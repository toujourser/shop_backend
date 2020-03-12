package models

import "database/sql"

// 权限表
type Permission struct {
	PsId    int64  `gorm:"primary_key;AUTO_INCREMENT" json:"ps_id"`
	PsName  string `gorm:"size(20)" json:"ps_name"`          // 权限名称
	PsPid   int    `gorm:"size(6)" json:"ps_pid"`            // 父id
	PsC     string `gorm:"size(32);default('')" json:"ps_c"` // 控制器
	PsA     string `gorm:"size(32);default('')" json:"ps_a"` // 操作方法
	PsLevel string `gorm:"default(0)" json:"ps_level"`       // 权限等级
}

// 权限API
type PermissionApi struct {
	Id           int64          `gorm:"primary_key" json:"id"`
	PermissionId int64          `json:"ps_id"`
	PsApiService sql.NullString `gorm:"size(255);NULL" json:"ps_api_service"`
	PsApiAction  sql.NullString `gorm:"size(255);NULL" json:"ps_api_action"`
	PsApiPath    sql.NullString `gorm:"size(255);NULL" json:"ps_api_path"`
	PsApiOrder   int            `gorm:"NULL" json:"ps_api_order"`
}
