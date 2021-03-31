package repo

import (
	"github.com/gofrs/uuid"
	"os"
	"path"
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
	dir := path.Join(r.dir, "temp")
	os.MkdirAll(dir, 0666)
	return dir
}

func (r *Repository) GetCacheDir() string {
	dir := path.Join(r.dir, "cache")
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
