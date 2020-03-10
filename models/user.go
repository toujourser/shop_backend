package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 会员表
type User struct {
	UserId        int64  `orm:"pk" json:"user_id"`
	Username      string `orm:"size(128)" json:"username"`             // 登录名
	QqOpenId      string `orm:"NULL;size(32)" json:"qq_open_id"`       // qq官方唯一编号信息
	Password      string `orm:"size(64)" json:"password"`              // 登录密码
	UserEmail     string `orm:"size(64)" json:"user_email"`            // 邮箱
	UserEmailCode string `orm:"size(13)" json:"user_email_code"`       // 新用户注册邮件激活唯一校验码
	IsActive      string `orm:"default(否)" json:"is_active"`           // 新用户是否已经通过邮箱激活帐号
	UserSex       string `orm:"default(男)" json:"user_sex"`            // 性别
	UserQq        string `orm:"size(32)" json:"user_qq"`               // qq
	UserTel       string `orm:"size(32)" json:"user_tel"`              // 手机
	UserXueli     string `orm:"default(本科)" json:"user_xueli"`        // 学历
	UserHobby     string `orm:"size(32)"json:"user_hobby"`             // 爱好
	UserIntroduce string `orm:"type(text);NULL" json:"user_introduce"` // 简介
	CreateTime    int    `orm:"size(17)" json:"create_time"`
	UpdateTime    int    `orm:"size(17)" json:"update_time"`
}

type UserCart struct {
	CartId     int64     `orm:"pk" json:"cart_id"`
	User       *User     `orm:"rel(fk)" json:"user_id"`
	CartInfo   string    `orm:"type(text)" json:"cart_info"` // 购物车详情信息，二维数组序列化信息
	CreateAt   time.Time `orm:"NULL;auto_now_add;type(datetime)" json:"create_at"`
	UpdateAt   time.Time `orm:"NULL;type(datetime)" json:"update_at"`
	DeleteTime time.Time `orm:"NULL"json:"delete_time"`
}

func decryptAES(enPwd, pwd string) bool {
	// 密码验证
	err := bcrypt.CompareHashAndPassword([]byte(enPwd), []byte(pwd))
	if err != nil {
		return true
	} else {
		return false
	}
}
