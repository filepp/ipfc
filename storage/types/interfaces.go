package types

import (
	"context"
	"github.com/ipfs/go-cid"
)

type Storage interface {
	AddFile(ctx context.Context, filePath string) (fileCid cid.Cid, err error)

	RetrieveFile(ctx context.Context, fileCid cid.Cid, outputPath string) error
}

type MinerSelectorFilter func(context.Context, *FileInfo) []string

type MinerSelector interface {
	GetMiners(ctx context.Context, fileInfo *FileInfo, count int, filter MinerSelectorFilter) ([]string, error)
}
