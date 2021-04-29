package reward

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	logging "github.com/ipfs/go-log/v2"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/eth"
	"ipfc/utils/encoding"
	"math"
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
	log.Info("createMiners")

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
	log.Infof("accountCount： %v", accountCount)

	for i := int64(index) + 1; i < accountCount; i++ {
		address, minerId, err := m.contract.GetAccount(i)
		if err != nil {
			log.Errorf("failed to get account:%v", err)
			return
		}
		log.Infof("miner： %v: %v", address, minerId)

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
		ticker := time.NewTicker(time.Minute)
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

	decimals, err := m.contract.GetDecimals()
	if err != nil {
		log.Errorf("failed to get decimals")
		return
	}

	for _, miner := range miners {
		// 最近一个小时在线的矿工
		if miner.LastActiveAt+60*60 > tNow {
			addr := common.HexToAddress(miner.Address)
			if !m.checkBalance(addr, decimals) {
				continue
			}
			indexes = append(indexes, miner.Idx)
		}
	}

	if len(indexes) > 0 {
		log.Infof("Approves: %v", indexes)
		data := encoding.EncodeIndex(indexes)
		tx, err := m.contract.Approves(data, len(indexes))
		if err != nil {
			log.Errorf("failed to approve")
			return
		}
		log.Infof("approve tx: %v", tx.Hash().String())
	}
}

func (m *Cron) checkBalance(addr common.Address, decimals uint8) bool {
	balance, err := m.contract.GetBalanceOf(addr)
	if err != nil {
		log.Errorf("failed to get balance: %v", err)
		return false
	}

	// 10 FC
	if float64(balance.Int64()) < math.Pow(10, float64(decimals))*10 {
		log.Errorf("no enough FC")
		return false
	}
	return true
}
