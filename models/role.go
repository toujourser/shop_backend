package models

// 用户角色表
type Role struct {
	RoleId   int64  `orm:"pk" json:"role_id"`
	RoleName string `orm:"size(20)" json:"role_name"`
	PsIds    string `orm:"size(512)" json:"ps_ids"`      // 权限ids,1,2,5
	PsCa     string `orm:"type(text);NULL" json:"ps_ca"` // 控制器-操作,控制器-操作,控制器-操作
	RoleDesc string `orm:"type(text);NULL" json:"role_desc"`
}
