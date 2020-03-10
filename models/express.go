package models

// 快递表
type Express struct {
	ExpressId  int64  `gorm:"primary_key" json:"express_id"`
	OrderId    uint64 `json:"order_id"`
	ExpressCom string `gorm:"size:32;NULL" json:"express_com"` // 订单快递公司名称
	ExpressNu  string `gorm:"size:32;NULL" json:"express_nu"`  // 快递单编号
	CreateTime int    `gorm:"size:11" json:"create_time"`
	UpdateTime int    `gorm:"size:11" json:"update_time"`
	//Order      *Order `gorm:"rel(fk)" json:"order_id"` // 订单id
}
