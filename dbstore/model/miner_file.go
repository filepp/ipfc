package model

type MinerFile struct {
	Id      int    `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	MinerId string `gorm:"index:idx_miner_id"`
	FileCid string `gorm:"index:idx_file_cid"`
	State   int    // 时空证明结果 0-正常， 1-失败
}
