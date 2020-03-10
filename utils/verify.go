package utils

import "regexp"

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyMobile(mobile string) bool {
	pattern := "^(0|86|17951)?(13[0-5]|15[0-9]|17[678]|18[0-9]|14[5-7])[0-9]{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}
