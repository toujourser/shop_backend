package models

// 管理员表
type Manager struct {
	MgId     int64  `gorm:"primary_key;AUTO_INCREMENT" json:"mg_id"`
	MgName   string `gorm:"size(32)" json:"mg_name"`
	MgPwd    string `gorm:"size(64)" json:"mg_pwd"`
	MgTime   int    `gorm:"size(10)" json:"mg_time"` // 注册时间
	MgMobile string `gorm:"size(11)" json:"mg_mobile"`
	MgEmail  string `gorm:"size(64)" json:"mg_email"`
	MgState  int    `gorm:"default(1)" json:"mg_state"` // 1：表示启用 0:表示禁用
	//Role     *Role  `gorm:"foreignkey:role_id" json:"role_id"` // 角色id
	RoleId   int64  `json:"role_id"`                 // 角色id
}
