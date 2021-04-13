package ipfs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/miner/proto"
	logging "github.com/ipfs/go-log/v2"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/storage/types"
	"ipfc/subpub"
	"os"
	"strings"
	"sync"
	"time"
)

var log = logging.Logger("ipfs")

type Storage struct {
	sync.RWMutex
	ipfsApi    *httpapi.HttpApi
	db         *ds.DbStore
	subscriber *subpub.Subscriber
	replicas   int
}

func NewStorage(peerId, addr string, replicas int, db *ds.DbStore) (*Storage, error) {
	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	ipfsApi, err := httpapi.NewApi(maddr)
	if err != nil {
		return nil, err
	}
	subscriber := subpub.NewSubscriber(peerId, ipfsApi, NewV1Handler(ipfsApi))
	subscriber.Subscribe()

	return &Storage{
		ipfsApi:    ipfsApi,
		db:         db,
		subscriber: subscriber,
		replicas:   replicas,
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
	log.Infof("Add file:%v, cid:%v", filePath, resolved.Cid())

	mfile := &model.File{
		Cid:       resolved.Cid().String(),
		State:     0,
		CreatedAt: time.Now().Unix(),
	}
	if s.db.FileExist(mfile.Cid) {
		return resolved.Cid(), nil
	}
	err = s.db.CreateFile(mfile)
	if err != nil {
		return resolved.Root(), err
	}
	s.syncToOtherPeers(ctx, resolved.Cid())

	return resolved.Cid(), nil
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

	//todo: 根据节点的容量等判断
	selector := NewPeerSelector()
	peers, _ = selector.GetPeers(ctx, s.filterPeers(s.allMiners(), peers), s.replicas)
	var mfiles []model.MinerFile
	for _, peer := range peers {
		err := s.publishFetchMessage(ctx, peer.ID(), fileCid)
		if err != nil {
			continue
		}
		mfiles = append(mfiles, model.MinerFile{
			MinerId: peer.ID().String(),
			FileCid: fileCid.String(),
			State:   0,
		})
	}
	if len(mfiles) > 0 {
		err = s.db.CreateMinerFiles(mfiles)
	}
	return err
}

func (s *Storage) publishFetchMessage(ctx context.Context, id peer.ID, fileCid cid.Cid) error {
	topic := proto.V1Topic(id.String())
	msg := proto.Message{
		Type:  proto.MsgFetchFile,
		Nonce: s.genNonce(),
		Data: proto.FetchFileReq{
			Cid: fileCid,
		},
	}
	data, _ := msg.EncodeMessage()
	err := s.ipfsApi.PubSub().Publish(ctx, topic, data)
	if err != nil {
		log.Errorf("failed to publish: %v", err)
		return err
	}
	return err
}

func (s *Storage) genNonce() string {
	nonce, _ := uuid.NewV4()
	return strings.ReplaceAll(nonce.String(), "-", "")
}

func (s *Storage) allMiners() map[string]*model.Miner {
	//todo: 缓存优化
	list, _ := s.db.GetAllMiners(-1, -1)
	miners := make(map[string]*model.Miner)
	for _, v := range list {
		miners[v.Id] = v
	}
	return miners
}

func (s *Storage) filterPeers(miners map[string]*model.Miner, peers []iface.ConnectionInfo) []iface.ConnectionInfo {
	list := make([]iface.ConnectionInfo, 0)
	for _, v := range peers {
		if _, ok := miners[v.ID().String()]; ok {
			list = append(list, v)
		}
	}
	return list
}

var _ types.Storage = &Storage{}
