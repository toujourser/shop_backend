package common

import (
	"github.com/kataras/iris/v12"
	"vue_shop/utils"
)

type Common struct {
	Ctx iris.Context
}

func (c *Common) ReturnJson(status int, msg string, args ...interface{}) {
	result := map[string]interface{}{}
	data := map[string]interface{}{}
	meta := map[string]interface{}{}

	meta["status"] = status
	meta["msg"] = msg
	result["meta"] = meta

	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case map[string]interface{}:
			for k, v := range arg.(map[string]interface{}) {
				data[k] = v
			}
		case string:
			key = arg.(string)
		default:
			data[key] = arg
		}

	}

	result["data"] = data

	c.Ctx.JSON(result)
	c.Ctx.StopExecution()
	return
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	//result["code"] = utils.RCODE_OK
	//result["message"] = utils.RecodeText(utils.RCODE_OK)

	if args != nil{
		result["data"] = args[0]
	}

	meta := map[string]interface{}{}
	meta["status"] = utils.RCODE_OK
	meta["msg"] = utils.RecodeText(utils.RCODE_OK)
	result["meta"] = meta

	//data := map[string]interface{}{}
	//key := ""
	//for _, arg := range args {
	//	switch arg.(type) {
	//	case string:
	//		key = arg.(string)
	//	default:
	//		data[key] = arg
	//	}
	//}
	//result["data"] = data


	this.Ctx.JSON(result)
	this.Ctx.StopExecution()
	return
}
