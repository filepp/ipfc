package lotus

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	logging "github.com/ipfs/go-log/v2"
	"ipfc/storage/types"
)

var log = logging.Logger("lotus")

type Storage struct {
	node   api.FullNode
	closer jsonrpc.ClientCloser
}

func NewStorage(apiAddr, token string) (*Storage, error) {
	node, closer, err := GetFullNodeAPIUsingCredentials(context.Background(), apiAddr, token)
	if err != nil {
		return nil, err
	}
	return &Storage{
		node:   node,
		closer: closer,
	}, nil
}

func (s *Storage) AddFile(ctx context.Context, filePath string) (cid string, err error) {
	return
}

func (s *Storage) RetrieveFile(ctx context.Context, cid, outputPath string) error {
	return nil
}

var _ types.Storage = &Storage{}