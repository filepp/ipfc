package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/liyue201/golib/xreflect"
	"github.com/liyue201/golib/xsignal"
	"ipfc/config"
	"ipfc/dbstore/ds"
	"ipfc/inspection"
	"ipfc/utils/xgorm"
)

var log = logging.Logger("main")

func main() {
	conf := config.AppConfig
	log.Infof("conf: %v", conf)

	dbConf := xgorm.DefaultConfig()
	xreflect.StructCopy(&conf.Mysql, dbConf)
	db := dbConf.Build()
	defer db.Close()

	dbStore := ds.NewDbStore(db)
	inspector, err := inspection.NewInspector(conf.Ipfs.PeerId, conf.Ipfs.ApiAddr, dbStore)
	if err != nil {
		log.Errorf("failed to new ipfs inspector: %v", err)
		return
	}
	inspector.Run()
	defer inspector.Stop()

	xsignal.Wait()
}
