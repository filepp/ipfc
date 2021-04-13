package model

type MinerFile struct {
	Id      string `gorm:"column:id; PRIMARY_KEY"`
	MinerId string `gorm:"index:idx_miner_id"`
	FileCid string `gorm:"index:idx_file_cid"`
	State   int
}
