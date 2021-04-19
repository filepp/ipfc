package model

// 时空证明结果
type WindowPost struct {
	Id        int `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	MinerId   string
	Request   string  // 请求
	Result    string  // 返回
	PassRate  float64 // 通过率
	CreatedAt int64   // 创建时间
}
