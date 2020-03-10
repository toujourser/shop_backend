package models

// 类型表
type Type struct {
	TypeId     int64  `orm:"pk" json:"type_id"`
	TypeName   string `orm:"size(32)" json:"type_name"`// 类型名称
	DeleteTime int    `orm:"size(17);NULL" json:"delete_time"`
}
