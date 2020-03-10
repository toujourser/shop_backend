package models

import "time"

type Report_1 struct {
	Id           int64     `orm:"pk" json:"id"`
	Rp1UserCount string    `orm:"size(8);NULL" json:"rp1_user_count"` // 用户数
	Rp1Area      string    `orm:"NULL" json:"rp1_area"`               // 地区
	Rp1Date      time.Time `orm:"NULL" json:"rp1_date"`
}

type Report_2 struct {
	Id       int64     `orm:"pk" json:"id"`
	Rp2Page  string    `orm:"size(128);NULL" json:"rp2_page"`
	Rp2Count int       `orm:"NULL" json:"rp2_count"`
	Rp2Date  time.Time `orm:"NULL" json:"rp2_date"`
}

type Report_3 struct {
	Id       int64     `orm:"pk" json:"id"`
	Rp3Src   string    `orm:"size(127);NULL" json:"rp3_src"` // 用户来源
	Rp3Count int       `orm:"NULL" json:"rp3_count"`         // 数量
	Rp3Date  time.Time `orm:"NULL" json:"rp3_date"`
}
