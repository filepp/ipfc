package types

type Storage interface {
	AddFile(filePath string) (cid string, err error)

	RetrieveFile(cid, outputPath string) error
}
