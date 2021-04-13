package model

type Miner struct {
	Id           string `gorm:"column:id; PRIMARY_KEY"` //  ipfs peerId
	Address      string // 钱包地址
	Role         int    // 角色 0-中心节点 1-边沿节点
	State        int    // 状态 0-在线 1-不在线
	CreatedAt    uint64 // 创建时间
	LastActiveAt uint64 // 最后在线时间
}

// 节点角色
const (
	// 中心节点
	RoleMainNode = 0

	// 边沿节点
	RoleEdgeNode = 1
)
