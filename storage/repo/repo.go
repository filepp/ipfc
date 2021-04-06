package repo

import (
	"github.com/gofrs/uuid"
	"os"
	"path"
)

const (
	tempDir      = "temp"
	cacheDir     = "cache"
	dataStoreDir = "datastore"
)

type Repository struct {
	dir string
}

func NewRepository(dir string) *Repository {
	return &Repository{
		dir: dir,
	}
}

func (r *Repository) GetTempDir() string {
	dir := path.Join(r.dir, tempDir)
	os.MkdirAll(dir, 0666)
	return dir
}

func (r *Repository) GetCacheDir() string {
	dir := path.Join(r.dir, cacheDir)
	os.MkdirAll(dir, 0666)
	return dir
}

func (r *Repository) GetDataStoreDir() string {
	dir := path.Join(r.dir, dataStoreDir)
	os.MkdirAll(dir, 0666)
	return dir
}

func (r *Repository) GetTempFilePath() string {
	fileUuid, _ := uuid.NewV4()
	filepath := path.Join(r.GetTempDir(), fileUuid.String())
	return filepath
}

func (r *Repository) GetCacheFilePath(fileCid string) string {
	filepath := path.Join(r.GetCacheDir(), fileCid)
	return filepath
}
