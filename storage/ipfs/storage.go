package ipfs

import (
	"context"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs-http-client"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/storage/types"
	"os"
)

var log = logging.Logger("ipfs")

type Storage struct {
	ipfsApi *httpapi.HttpApi
}

func NewStorage(addr string) (*Storage, error) {
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

var _ types.Storage = &Storage{}
