package mananer

import (
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

func (m *Manager) AddFile(filePath string) (cid string, err error) {
	cid, err = m.hotStorage.AddFile(filePath)
	if err != nil {
		return "", err
	}
	return m.coldStorage.AddFile(filePath)
}

func (m *Manager) RetrieveFile(cid, outputPath string) error {
	err := m.hotStorage.RetrieveFile(cid, outputPath)
	if err != nil {
		log.Warnf("failed to retrieve file from hot storage: %v", err.Error())
		return m.coldStorage.RetrieveFile(cid, outputPath)
	}
	return nil
}
