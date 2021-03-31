package ipfs

import (
	"github.com/ipfs/go-ipfs-http-client"
	logging "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"
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
