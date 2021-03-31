package lotus

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	logging "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"net/http"
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

func GetFullNodeAPIUsingCredentials(ctx context.Context, apiAddr, token string) (api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(apiAddr)
	if err != nil {
		return nil, nil, err
	}

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {
		return nil, nil, err
	}

	return NewFullNodeRPC(ctx, apiURI(addr), apiHeaders(token))
}

func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}

func apiHeaders(token string) http.Header {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	return headers
}

func NewFullNodeRPC(ctx context.Context, addr string, requestHeader http.Header) (api.FullNode, jsonrpc.ClientCloser, error) {
	var res apistruct.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin",
		[]interface{}{
			&res.CommonStruct.Internal,
			&res.Internal,
		}, requestHeader)

	return &res, closer, err
}

