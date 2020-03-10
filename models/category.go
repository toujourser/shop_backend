package models

// 分类表
type Category struct {
	CatId      int64  `orm:"pk" json:"cat_id"`
	CatName    string `orm:"size(255);NULL" json:"cat_name"`
	CatPid     int    `orm:"NULL" json:"cat_pid"` // 分类父ID
	CatLevel   int    `orm:"NULL" json:"cat_level"` // 分类层级 0: 顶级 1:二级 2:三级
	CatDeleted int8   `orm:"default(0);NULL" json:"cat_deleted"` // 是否删除 1为删除
	CatIcon    string `orm:"size(255);NULL" json:"cat_icon"`
	CatSrc     string `orm:"type(text);NULL" json:"cat_src"`
}
