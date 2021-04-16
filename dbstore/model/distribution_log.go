package model

type DistributionLog struct {
	Id        int `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	Amount    string
	CreatedAt int64
}
