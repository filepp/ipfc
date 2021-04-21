package reward

import (
	"context"
	logging "github.com/ipfs/go-log/v2"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/eth"
	"ipfc/utils/encoding"
	"sync"
	"time"
)

var log = logging.Logger("reword")

type Config struct {
	Amount int64 `yaml:"amount"`
}

type Cron struct {
	conf     Config
	contract *eth.Contract
	store    *ds.DbStore
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

func NewCron(store *ds.DbStore, contract *eth.Contract, conf Config) Cron {
	return Cron{
		conf:     conf,
		store:    store,
		contract: contract,
	}
}

func (m *Cron) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	m.runCreateMiner(ctx)
	m.runDistributeToken(ctx)
}

func (m *Cron) Stop() {
	m.cancel()
	m.wg.Wait()
}

func (m *Cron) runCreateMiner(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer func() {
			m.wg.Done()
		}()
		m.createMiners(ctx)
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				m.createMiners(ctx)
			}
		}
	}()
}

func (m *Cron) createMiners(ctx context.Context) {
	index, err := m.store.GetMaxMinerIndex()
	if err != nil {
		log.Errorf("failed to get miner count:%v", err)
		return
	}
	accountCount, err := m.contract.GetAccountCount()
	if err != nil {
		log.Errorf("failed to get miner count:%v", err)
		return
	}
	for i := int64(index); i < accountCount; i++ {
		address, minerId, err := m.contract.GetAccount(i)
		if err != nil {
			log.Errorf("failed to get account:%v", err)
			return
		}
		if m.store.MinerExist(minerId) {
			continue
		}
		miner := &model.Miner{
			Id:        minerId,
			Address:   address,
			Role:      model.RoleEdgeNode,
			Idx:       int(i),
			CreatedAt: time.Now().Unix(),
		}
		err = m.store.CreateMiner(miner)
		if err != nil {
			log.Errorf("failed to create miner:%v", err)
			return
		}
	}
}

func (m *Cron) runDistributeToken(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer func() {
			m.wg.Done()
		}()
		//ticker := time.NewTicker(time.Minute * 10)
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				m.distributeToken(ctx)
			}
		}
	}()
}

func (m *Cron) distributeToken(ctx context.Context) {
	miners, err := m.store.GetAllMiners(-1, -1)
	if err != nil {
		log.Errorf("failed to get miners")
		return
	}
	var indexes []int
	tNow := time.Now().Unix()

	for _, miner := range miners {
		// 最近一个小时在线的矿工
		if miner.LastActiveAt+60*60 > tNow {
			indexes = append(indexes, miner.Idx)
		}
	}
	if len(indexes) > 0 {
		data := encoding.EncodeIndex(indexes)
		tx, err := m.contract.Approves(data, len(indexes))
		if err != nil {
			log.Errorf("failed to approve")
			return
		}
		log.Infof("approve tx: %v", tx.Hash().String())
	}
}
