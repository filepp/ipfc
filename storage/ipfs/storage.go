package ipfs

import (
	"context"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/miner/proto"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/dbstore/ds"
	"ipfc/storage/types"
	"os"
)

var log = logging.Logger("ipfs")

type Storage struct {
	ipfsApi *httpapi.HttpApi
	db      *ds.DbStore
}

func NewStorage(addr string, db *ds.DbStore) (*Storage, error) {
	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	httpApi, err := httpapi.NewApi(maddr)
	if err != nil {
		return nil, err
	}
	return &Storage{
		ipfsApi: httpApi,
		db:      db,
	}, nil
}

func (s *Storage) AddFile(ctx context.Context, filePath string) (cid.Cid, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("%v, filePath=%v", err.Error(), filePath)
		return cid.Undef, err
	}
	fileNode := files.NewReaderFile(file)
	resolved, err := s.ipfsApi.Unixfs().Add(ctx, fileNode, options.Unixfs.CidVersion(1), options.Unixfs.Pin(true))
	if err != nil {
		log.Errorf("%v", err.Error())
		return cid.Undef, err
	}
	log.Infof("Add file:%v, cid:%v, %v", filePath, resolved.Cid())
	s.syncToOtherPeers(ctx, resolved.Cid())
	return resolved.Root(), nil
}

func (s *Storage) RetrieveFile(ctx context.Context, fileCid cid.Cid, outputPath string) error {
	cidPath := path.New("/ipfs/" + fileCid.String())
	fileNode, err := s.ipfsApi.Unixfs().Get(ctx, cidPath)
	if err != nil {
		log.Errorf("%v", err.Error())
		return err
	}
	defer fileNode.Close()

	err = files.WriteTo(fileNode, outputPath)
	if err != nil {
		log.Errorf("%v", err.Error())
		return err
	}
	return nil
}

// 备份数据
// 给其他节点发消息, 其他节点收到消息后同步数据
func (s *Storage) syncToOtherPeers(ctx context.Context, fileCid cid.Cid) error {
	peers, err := s.ipfsApi.Swarm().Peers(ctx)
	if err != nil {
		log.Errorf("failed to get peers: %v", err.Error())
		return err
	}

	replicas := 3
	selector := NewPeerSelector()
	peers, _ = selector.GetPeers(ctx, peers, replicas)
	for _, peer := range peers {
		s.publishFetchMessage(ctx, peer.ID(), fileCid)
	}
	return nil
}

func (s *Storage) publishFetchMessage(ctx context.Context, id peer.ID, fileCid cid.Cid) error {
	topic := proto.V1Topic(id.String())
	msg := proto.Message{
		Type: proto.MsgFetchFile,
		Data: proto.FetchFile{
			Cid: fileCid,
		},
	}
	data, _ := msg.EncodeMessage()
	err := s.ipfsApi.PubSub().Publish(ctx, topic, data)
	if err != nil {
		return err
	}
	// TODO： 处理应答消息
	return err
}

var _ types.Storage = &Storage{}
