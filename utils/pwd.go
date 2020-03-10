package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptAES(src []byte) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		panic("密码加密错误")
	}

	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePW
}

func CompareAES(enPwd, pwd string) bool {
	// 密码验证
	err := bcrypt.CompareHashAndPassword([]byte(enPwd), []byte(pwd))
	if err != nil {
		return false
	} else {
		return true
	}
}
