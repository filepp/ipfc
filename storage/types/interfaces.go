package types

import "context"

type Storage interface {
	AddFile(ctx context.Context, filePath string) (cid string, err error)

	RetrieveFile(ctx context.Context, cid, outputPath string) error
}
