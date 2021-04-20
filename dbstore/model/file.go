package model

type File struct {
	Cid        string `gorm:"column:cid; PRIMARY_KEY"`
	CidInLotus string //文件在lotus中的cid
	Size       int64
	State      int // 0-正常 1-异常

	CreatedAt  int64
}
