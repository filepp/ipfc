package model

type File struct {
	Cid       string `gorm:"column:cid; PRIMARY_KEY"`
	State     int    // 0-正常 1-异常
	CreatedAt int64
}
