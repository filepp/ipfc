package inspection

import (
	"context"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/miner/proto"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/dbstore/ds"
	"ipfc/subpub"
	"sync"
	"time"
)

type Inspector struct {
	sync.RWMutex
	ipfsApi    *httpapi.HttpApi
	store      *ds.DbStore
	subscriber *subpub.Subscriber
	wg         sync.WaitGroup
	cancel     context.CancelFunc
}

func NewInspector(peerId, addr string, store *ds.DbStore) (*Inspector, error) {
	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	ipfsApi, err := httpapi.NewApi(maddr)
	if err != nil {
		return nil, err
	}
	handler := NewV1Handler(ipfsApi, store)

	subscriber := subpub.NewSubscriber(ipfsApi)
	subscriber.Subscribe(proto.V1MinerHeartBeatTopic(), handler.Handle)
	subscriber.Subscribe(proto.V1ExternalTopic(peerId), handler.Handle)
	return &Inspector{
		ipfsApi:    ipfsApi,
		store:      store,
		subscriber: subscriber,
	}, nil
}

func (i *Inspector) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	i.cancel = cancel
	i.wg.Add(1)
	go func() {
		defer func() {
			i.wg.Done()
		}()
		ticker := time.NewTicker(time.Hour * 2)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				i.doInspect(ctx)
			}
		}
	}()
}

func (i *Inspector) Stop() {
	i.cancel()
	i.wg.Wait()
}

func (i *Inspector) doInspect(ctx context.Context) {

}
