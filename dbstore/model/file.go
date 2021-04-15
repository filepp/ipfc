package model

type File struct {
	Cid       string `gorm:"column:cid; PRIMARY_KEY"`
	Size      int64
	State     int // 0-正常 1-异常
	CreatedAt int64
}
