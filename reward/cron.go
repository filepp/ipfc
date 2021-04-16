package reward

import (
	"context"
	"github.com/ethereum/go-ethereum/common/math"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jinzhu/gorm"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/eth"
	"math/big"
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
	m.wg.Add(1)
	go func() {
		defer func() {
			m.wg.Done()
		}()
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if time.Now().Hour() == 0 {
					disLog, err := m.store.GetLastDistributionLog()
					if err != nil && err != gorm.ErrRecordNotFound {
						continue
					}
					if disLog.CreatedAt+60*60*24 < time.Now().Unix() {
						m.distributeToken(ctx)
					}
				}
			}
		}
	}()
}

func (m *Cron) Stop() {
	m.cancel()
	m.wg.Wait()
}

func (m *Cron) distributeToken(ctx context.Context) {
	miners, err := m.store.GetAllMiners(-1, -1)
	if err != nil {
		log.Errorf("failed to get miners")
		return
	}
	var minerList []*model.Miner
	tNow := time.Now().Unix()

	for _, miner := range miners {
		// 最近一个小时在线的矿工
		if miner.LastActiveAt+60*60 > tNow {
			minerList = append(minerList, miner)
		}
	}
	reward := big.NewInt(0).Div(m.rewardPerDay(), big.NewInt(int64(len(minerList))))
	var (
		addresses []string
		rewards   []*big.Int
	)
	batch := 10
	for _, miner := range minerList {
		select {
		case <-ctx.Done():
			return
		default:
		}
		addresses = append(addresses, miner.Address)
		rewards = append(rewards, reward)
		if len(addresses) == batch {
			m.sendApproves(addresses, rewards)
			addresses = make([]string, 0)
			rewards = make([]*big.Int, 0)
		}
	}
	if len(addresses) > 0 {
		m.sendApproves(addresses, rewards)
	}
	m.store.CreateDistributionLog(&model.DistributionLog{
		Amount:    reward.String(),
		CreatedAt: time.Now().Unix(),
	})
}

func (m *Cron) sendApproves(addresses []string, rewards []*big.Int) (err error) {
	for i := 0; i < 3; i++ {
		_, err = m.contract.Approves(addresses, rewards)
		if err != nil {
			log.Errorf("failed to approve: %v", err.Error())
			time.Sleep(time.Minute)
			continue
		}
		return nil
	}
	return err
}

func (m *Cron) rewardPerDay() *big.Int {
	decimals, err := m.contract.GetDecimals()
	if err != nil {
		decimals = 0
	}
	mm := math.BigPow(int64(10), int64(decimals))
	return big.NewInt(0).Mul(big.NewInt(m.conf.Amount), mm)
}
