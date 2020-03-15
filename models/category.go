package models

// 分类表
type Category struct {
	CatId      int64  `gorm:"primary_key;AUTO_INCREMENT" json:"cat_id"`
	CatName    string `gorm:"size(255);NULL" json:"cat_name"`
	CatPid     int    `gorm:"NULL" json:"cat_pid"`                // 分类父ID
	CatLevel   int    `gorm:"NULL" json:"cat_level"`              // 分类层级 0: 顶级 1:二级 2:三级
	CatDeleted int    `gorm:"default(0);NULL" json:"cat_deleted"` // 是否删除 1为删除
	CatIcon    string `gorm:"size(255);NULL" json:"cat_icon"`
	CatSrc     string `gorm:"type(text);NULL" json:"cat_src"`
}
