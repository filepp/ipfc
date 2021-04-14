package subpub

import (
	"context"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/miner/proto"
	logging "github.com/ipfs/go-log/v2"
	"runtime/debug"
	"time"
)

var log = logging.Logger("subpub")

type Subscriber struct {
	ipfsApi *httpapi.HttpApi
}

func NewSubscriber(ipfsApi *httpapi.HttpApi) *Subscriber {
	return &Subscriber{
		ipfsApi: ipfsApi,
	}
}

func (s *Subscriber) Subscribe(topic string, handler HandleFunc) error {
	sub, err := s.ipfsApi.PubSub().Subscribe(context.TODO(), topic)
	if err != nil {
		log.Errorf("failed to subscribe: %v", err)
		return err
	}
	log.Infof("subscribe: %v", topic)

	go func() {
		for {
			pmsg, err := sub.Next(context.Background())
			if err != nil {
				log.Errorf("failed get message: %v", err)
				time.Sleep(time.Second)
				continue
			}
			log.Infof("received message from %v", pmsg.From().String())
			go func() {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("%v", string(debug.Stack()))
					}
				}()
				msg, err := proto.DecodeMessage(pmsg.Data())
				if err != nil {
					log.Errorf("failed to decode message: %v", err)
					return
				}
				err = handler(context.TODO(), pmsg.From(), &msg)
				if err != nil {
					log.Errorf("failed to handler message: %v", err)
				}
			}()
		}
	}()
	return nil
}
