package models

// 属性表
type Attribute struct {
	AttrId     int64  `gorm:"pk" json:"attr_id"`
	AttrName   string `gorm:"size:32" json:"attr_name"`
	CategoryId uint64 `json:"category_id"`
	AttrSel    string `gorm:"default:only" json:"attr_sel"`     // only:输入框(唯一)  many:后台下拉列表/前台单选框
	AttrWrite  string `gorm:"default:manual" json:"attr_write"` // manual:手工录入  list:从列表选择
	AttrVals   string `gorm:"type:text;NULL" json:"attr_vals"`  // 可选值列表信息,例如颜色：白色,红色,绿色,多个可选值通过逗号分隔
	DeleteTime int    `gorm:"NULL" json:"delete_time"`
	//Category   *Category `gorm:"rel(fk)" json:"cat_id"`
}
