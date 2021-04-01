package lotus

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/ipfs/go-cid"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"net/http"
	"sort"
)

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

func getRetrievalOffers(ctx context.Context, node api.FullNode, payloadCid cid.Cid, pieceCid *cid.Cid, miners []string) []api.QueryOffer {

	// Ask each miner about costs and information about retrieving this data.
	var offers []api.QueryOffer
	for _, mi := range miners {
		a, err := address.NewFromString(mi)
		if err != nil {
			log.Infof("parsing miner address: %s", err)
		}
		qo, err := node.ClientMinerQueryOffer(ctx, a, payloadCid, pieceCid)
		if err != nil {
			log.Infof("asking miner %s query-offer failed: %s", a, err)
			continue
		}
		offers = append(offers, qo)
	}

	// Sort received options by price.
	sort.Slice(offers, func(a, b int) bool { return offers[a].MinPrice.LessThan(offers[b].MinPrice) })
	return offers
}