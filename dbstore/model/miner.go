package model

type Miner struct {
	Id           string `gorm:"column:id; PRIMARY_KEY"`
	Role         int
	State        int
	CreatedAt    uint64
	LastActiveAt uint64
}

// 节点角色
const (
	// 中心节点
	RoleMainNode = 0

	// 边沿节点
	RoleEdgeNode = 1
)
