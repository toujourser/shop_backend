package models

// 用户角色表
type Role struct {
	RoleId   int64  `gorm:"primary_key;AUTO_INCREMENT" json:"role_id"`
	RoleName string `gorm:"size(20)" json:"role_name"`
	PsIds    string `gorm:"size(512)" json:"ps_ids"`      // 权限ids,1,2,5
	PsCa     string `gorm:"type(text);NULL" json:"ps_ca"` // 控制器-操作,控制器-操作,控制器-操作
	RoleDesc string `gorm:"type(text);NULL" json:"role_desc"`
}
