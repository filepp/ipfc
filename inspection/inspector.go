package inspection

import (
	"context"
	"github.com/ipfs/go-cid"
	ds2 "github.com/ipfs/go-datastore"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/go-ipfs/miner/proto"
	logging "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/subpub"
	"ipfc/utils/xrand"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var log = logging.Logger("inspection")

const (
	InspectFileNum    = 2
	RandomPositionNum = 10
)

type Inspector struct {
	sync.RWMutex
	ipfsApi    *httpapi.HttpApi
	store      *ds.DbStore
	localStore ds2.Datastore
	subscriber *subpub.Subscriber
	wg         sync.WaitGroup
	cancel     context.CancelFunc
}

func NewInspector(peerId, addr string, store *ds.DbStore, localStore ds2.Datastore) (*Inspector, error) {
	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	ipfsApi, err := httpapi.NewApi(maddr)
	if err != nil {
		return nil, err
	}
	handler := NewV1Handler(ipfsApi, store, localStore)

	subscriber := subpub.NewSubscriber(ipfsApi)
	subscriber.Subscribe(proto.V1MinerHeartBeatTopic(), handler.Handle)
	subscriber.Subscribe(proto.V1InspectorTopic(peerId), handler.Handle)
	return &Inspector{
		ipfsApi:    ipfsApi,
		store:      store,
		localStore: localStore,
		subscriber: subscriber,
	}, nil
}

func (m *Inspector) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	m.wg.Add(1)
	go func() {
		defer func() {
			m.wg.Done()
		}()

		// 两个小时一次巡检，一天做12次, UTC时间0点为分割点
		slots := [12]bool{}
		preTwoHour := int(time.Now().Unix() / 60 / 60 / 2)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				twoHour := int(time.Now().Unix() / 60 / 60 / 2)
				if preTwoHour != twoHour && twoHour%12 == 0 {
					// 新的一天开始, 清零
					slots = [12]bool{}
				}
				if slots[twoHour%12] {
					continue
				}
				m.inspect(ctx)
				slots[twoHour%12] = true
			}
		}
	}()
}

func (m *Inspector) Stop() {
	m.cancel()
	m.wg.Wait()
}

func (m *Inspector) inspect(ctx context.Context) {
	miners, err := m.store.GetAllMiners(model.RoleEdgeNode, -1)
	if err != nil {
		log.Errorf("failed to get miners")
		return
	}
	for _, miner := range miners {
		select {
		case <-ctx.Done():
			return
		default:
		}
		err := m.inspectMiner(ctx, miner)
		if err != nil {
			return
		}
	}
}

func (m *Inspector) inspectMiner(ctx context.Context, miner *model.Miner) error {
	files, err := m.store.GetMinerFiles(miner.Id)
	if err != nil {
		log.Errorf("failed to get miner files")
		return err
	}
	if len(files) == 0 {
		return nil
	}
	files = selectFiles(files, InspectFileNum)

	req := proto.WindowPostReq{
		Items: make([]proto.WindowPostReqItem, 0),
	}
	for _, file := range files {
		if file.Size == 0 {
			continue
		}
		fileCid, err := cid.Decode(file.FileCid)
		if err != nil {
			log.Errorf("failed to decode cid: %v", err)
			continue
		}
		req.Items = append(req.Items, proto.WindowPostReqItem{
			FileCid:   fileCid,
			Positions: genRandomPositions(file.Size, RandomPositionNum),
		})
	}
	msgId := xrand.GenNonce()
	msg := proto.Message{
		Type:  proto.MsgWindowPost,
		Nonce: msgId,
		Data:  req,
	}
	data, _ := msg.EncodeMessage()
	err = m.ipfsApi.PubSub().Publish(ctx, proto.V1InternalTopic(miner.Id), data)
	if err != nil {
		log.Errorf("failed to publish: %v", err)
		return err
	}
	saveWindowPostMessage(m.localStore, miner.Id, msgId, req)
	return nil
}

func selectFiles(files []*ds.MinerFile, num int) []*ds.MinerFile {
	selected := make([]*ds.MinerFile, 0)
	if len(files) <= num {
		return files
	}
	mark := make(map[int]struct{})
	for i := 0; i < num; {
		pos := rand.Intn(len(files))
		if _, ok := mark[pos]; ok {
			continue
		}
		selected = append(selected, files[i])
		mark[pos] = struct{}{}
		i++
	}
	return selected
}

func genRandomPositions(size, num int64) []int64 {
	positions := make([]int64, 0)
	if size <= num {
		for i := int64(0); i < size; i++ {
			positions = append(positions, i)
		}
		return positions
	}
	mark := make(map[int64]struct{})
	for i := 0; i < int(num); {
		pos := rand.Int63n(size) % size
		if _, ok := mark[pos]; ok {
			continue
		}
		positions = append(positions, pos)
		mark[pos] = struct{}{}
		i++
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i] < positions[j]
	})
	return positions
}
