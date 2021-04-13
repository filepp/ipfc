package inspection

import (
	"context"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/dbstore/ds"
	"ipfc/subpub"
	"sync"
	"time"
)

type Inspector struct {
	sync.RWMutex
	ipfsApi    *httpapi.HttpApi
	db         *ds.DbStore
	subscriber *subpub.Subscriber
	wg         sync.WaitGroup
	cancel     context.CancelFunc
}

func NewInspector(peerId, addr string, db *ds.DbStore) (*Inspector, error) {
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

	return &Inspector{
		ipfsApi:    ipfsApi,
		db:         db,
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
