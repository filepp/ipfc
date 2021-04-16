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
	conf := AppConfig
	log.Infof("conf: %v", conf)

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
