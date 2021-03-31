package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/liyue201/golib/xsignal"
	"ipfc/rest"
	"ipfc/storage/ipfs"
	"ipfc/storage/lotus"
	"ipfc/storage/mananer"
)

var log = logging.Logger("main")

func main() {
	log.Infof("conf: %v", appConfig)

	ipfsStorage, err := ipfs.NewStorage(appConfig.Ipfs.ApiAddr)
	if err != nil {
		log.Errorf("failed to new ipfs storage: %v", err)
		return
	}
	lotusStorage, err := lotus.NewStorage(appConfig.Lotus.ApiAddr, appConfig.Lotus.Token)
	if err != nil {
		log.Errorf("failed to new lotus storage: %v", err)
		return
	}
	storage := mananer.NewManager(ipfsStorage, lotusStorage)
	server := rest.NewServer(appConfig.Http.ListenAddress, storage)
	server.Run()
	defer server.Stop()
	xsignal.Wait()
}
