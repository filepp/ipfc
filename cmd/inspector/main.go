package main

import (
	badger "github.com/ipfs/go-ds-badger2"
	logging "github.com/ipfs/go-log/v2"
	"github.com/liyue201/golib/xreflect"
	"github.com/liyue201/golib/xsignal"
	"ipfc/dbstore/ds"
	"ipfc/eth"
	"ipfc/inspection"
	"ipfc/reward"
	"ipfc/storage/repo"
	"ipfc/utils/xgorm"
)

var log = logging.Logger("main")

func main() {
	conf := AppConfig
	log.Infof("conf: %v", conf)

	dbConf := xgorm.DefaultConfig()
	xreflect.StructCopy(&conf.Mysql, dbConf)
	db := dbConf.Build()
	defer db.Close()

	repository := repo.NewRepository(conf.Repo.Dir)
	localStore, err := badger.NewDatastore(repository.GetDataStoreDir(), &badger.DefaultOptions)
	if err != nil {
		log.Errorf("failed to new datastore: %v", err)
		return
	}
	contract, err := eth.NewContract(conf.Eth)
	if err != nil {
		log.Errorf("failed to new contract: %v", err)
		return
	}

	dbStore := ds.NewDbStore(db)
	inspector, err := inspection.NewInspector(conf.Ipfs.PeerId, conf.Ipfs.ApiAddr, dbStore, localStore)
	if err != nil {
		log.Errorf("failed to new ipfs inspector: %v", err)
		return
	}
	inspector.Run()
	defer inspector.Stop()

	cron := reward.NewCron(dbStore, contract, conf.Reward)
	cron.Run()
	defer cron.Stop()

	xsignal.Wait()
}
