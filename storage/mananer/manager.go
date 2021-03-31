package mananer

import (
	"context"
	"github.com/prometheus/common/log"
	"ipfc/storage/types"
)

type Manager struct {
	hotStorage  types.Storage
	coldStorage types.Storage
}

func NewManager(hotStorage, coldStorage types.Storage) *Manager {
	return &Manager{
		hotStorage:  hotStorage,
		coldStorage: coldStorage,
	}
}

func (m *Manager) AddFile(ctx context.Context, filePath string) (cid string, err error) {
	cid, err = m.hotStorage.AddFile(ctx, filePath)
	if err != nil {
		return "", err
	}
	return m.coldStorage.AddFile(ctx, filePath)
}

func (m *Manager) RetrieveFile(ctx context.Context, cid, outputPath string) error {
	err := m.hotStorage.RetrieveFile(ctx, cid, outputPath)
	if err != nil {
		log.Warnf("failed to retrieve file from hot storage: %v", err.Error())
		return m.coldStorage.RetrieveFile(ctx, cid, outputPath)
	}
	return nil
}

var _ types.Storage = &Manager{}
