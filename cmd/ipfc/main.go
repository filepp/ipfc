package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/liyue201/golib/xreflect"
	"github.com/liyue201/golib/xsignal"
	"ipfc/api"
	"ipfc/dbstore/ds"
	"ipfc/storage/ipfs"
	"ipfc/storage/lotus"
	"ipfc/storage/mananer"
	"ipfc/storage/repo"
	"ipfc/utils/xgorm"
)

var log = logging.Logger("main")

func main() {
	log.Infof("conf: %v", appConfig)

	repository := repo.NewRepository(appConfig.Repo.Dir)
	dbConf := xgorm.DefaultConfig()
	xreflect.StructCopy(&appConfig.Mysql, dbConf)
	db := dbConf.Build()
	defer db.Close()
	dbStore := ds.NewDbStore(db)

	ipfsStorage, err := ipfs.NewStorage(appConfig.Ipfs.ApiAddr, dbStore)
	if err != nil {
		log.Errorf("failed to new ipfs storage: %v", err)
		return
	}
	lotusStorage, err := lotus.NewStorage(appConfig.Lotus.ApiAddr, appConfig.Lotus.Token, appConfig.Deal)
	if err != nil {
		log.Errorf("failed to new lotus storage: %v", err)
		return
	}
	storage := mananer.NewManager(ipfsStorage, lotusStorage, dbStore)
	server := api.NewServer(appConfig.Http.ListenAddress, storage, repository)
	server.Run()
	defer server.Stop()
	xsignal.Wait()
}
