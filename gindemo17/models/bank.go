package models

// foreignKey 外键，如果是表名称加上 Id，默认也可以不配置，如果不是我们需要通过 foreignKey 来配置
// references 表示的是主键，默认就是 Id，如果是 Id 的话可以不配置
type Bank struct {
	Id      int	// 主键
	Username   string
	Balance   int
}

func (Bank) TableName() string {
	return "bank"
}