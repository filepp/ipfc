package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/liyue201/golib/xreflect"
	"github.com/liyue201/golib/xsignal"
	"ipfc/api"
	"ipfc/dbstore/ds"
	"ipfc/eth"
	"ipfc/storage/ipfs"
	"ipfc/storage/lotus"
	"ipfc/storage/mananer"
	"ipfc/storage/repo"
	"ipfc/utils/xgorm"
	"math"
)

var log = logging.Logger("main")

func checkBalanceOfWallet() bool {
	contract, err := eth.NewContract(AppConfig.Eth)
	if err != nil {
		log.Errorf("failed to new contract: %v", err)
		return false
	}
	balance, err := contract.GetBalanceOfWallet()
	if err != nil {
		log.Errorf("failed to get balance: %v", err)
		return false
	}
	decimals, err := contract.GetDecimals()
	if err != nil {
		log.Errorf("failed to get decimals: %v", err)
		return false
	}

	if float64(balance.Int64()) < math.Pow(10, float64(decimals))*1_000_000_000 {
		log.Errorf("no enough FC")
		return false
	}
	return true
}

func main() {
	conf := AppConfig
	log.Infof("conf: %v", conf)
	if !checkBalanceOfWallet() {
		return
	}

	repository := repo.NewRepository(conf.Repo.Dir)
	dbConf := xgorm.DefaultConfig()
	xreflect.StructCopy(&conf.Mysql, dbConf)
	db := dbConf.Build()
	defer db.Close()
	dbStore := ds.NewDbStore(db)

	ipfsStorage, err := ipfs.NewStorage(conf.Ipfs.PeerId, conf.Ipfs.ApiAddr, conf.Ipfs.Replicas, dbStore)
	if err != nil {
		log.Errorf("failed to new ipfs storage: %v", err)
		return
	}
	lotusStorage, err := lotus.NewStorage(conf.Lotus.ApiAddr, conf.Lotus.Token, conf.Deal)
	if err != nil {
		log.Errorf("failed to new lotus storage: %v", err)
		return
	}
	storage := mananer.NewManager(ipfsStorage, lotusStorage, dbStore)
	server := api.NewServer(conf.Http.ListenAddress, storage, repository)
	server.Run()
	defer server.Stop()
	xsignal.Wait()
}
