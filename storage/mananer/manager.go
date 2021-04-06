package mananer

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/prometheus/common/log"
	"ipfc/storage/types"
)

type Manager struct {
	hotStorage  types.Storage
	coldStorage types.Storage
	datastore   datastore.TxnDatastore
}

func NewManager(hotStorage, coldStorage types.Storage, datastore datastore.TxnDatastore) *Manager {
	return &Manager{
		hotStorage:  hotStorage,
		coldStorage: coldStorage,
		datastore:   datastore,
	}
}

func (m *Manager) AddFile(ctx context.Context, filePath string) (fileCid cid.Cid, err error) {
	fileCid, err = m.hotStorage.AddFile(ctx, filePath)
	if err != nil {
		log.Infof("failed to add file to hot storage")
		return fileCid, err
	}

	_, err = m.coldStorage.AddFile(ctx, filePath)
	if err != nil {
		log.Infof("failed to add file to old storage")
		return fileCid, err
	}
	return fileCid, err
}

func (m *Manager) RetrieveFile(ctx context.Context, fileCid cid.Cid, outputPath string) error {
	err := m.hotStorage.RetrieveFile(ctx, fileCid, outputPath)
	if err != nil {
		log.Infof("failed to retrieve file from hot storage: %v", err.Error())
		err := m.coldStorage.RetrieveFile(ctx, fileCid, outputPath)
		if err != nil {
			return err
		}
		m.hotStorage.AddFile(ctx, outputPath)
	}
	return nil
}

var _ types.Storage = &Manager{}
