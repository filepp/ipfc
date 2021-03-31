package mananer

import "ipfc/storage/types"

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

func (m *Manager) SaveFile(filePath string) (cid string, err error) {
	return
}

func (m *Manager) RetrieveFile(cid, outputPath string) error {
	return nil
}
