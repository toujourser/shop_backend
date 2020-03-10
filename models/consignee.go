package models

// 收货人表
type Consignee struct {
	CgnId      int64  `gorm:"primary_key" json:"cgn_id"`
	UserId     uint64 `json:"user_id"`
	CgnName    string `gorm:"size(32)" json:"cgn_name"`
	CgnAddress string `gorm:"size(300);default('')" json:"cgn_address"`
	CgnTel     string `gorm:"size(20);default('')" json:"cgn_tel"`
	CgnCode    string `gorm:"default('')" json:"cgn_code"` // 邮编
	DeleteTime int    `gorm:"size(17);NULL" json:"delete_time"`
	//User       *User  `gorm:"rel(fk)" json:"user_id"`
}
