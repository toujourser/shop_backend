package main

import (
	"fmt"
	"vue_shop/utils"
)

func main() {
	//m := map[string][]map[string][]string{}
	//t1 := map[string][]string{
	//	"t1": []string{"aaa"},
	//}
	//
	//m["t"] = append(m["t"], t1)
	//fmt.Println(m)

	//argsTest("a", 1, "b", 2)


	println(utils.VerifyEmailFormat("tom@gmail.com"))
	println(utils.VerifyMobile("18861730375"))
}

func argsTest(args ...interface{}) {
	result := map[string]interface{}{}
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	println("---------------------")
	fmt.Printf("%+v", result)
}
