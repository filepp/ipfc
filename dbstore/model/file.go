package model

type File struct {
	Cid       string `gorm:"column:cid; PRIMARY_KEY"`
	State     int
	CreatedAt uint64
}
