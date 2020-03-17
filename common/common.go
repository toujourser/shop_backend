package common

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"vue_shop/utils"
)

type Common struct {
	Ctx iris.Context
}

func (c *Common) ParseParams(ctx iris.Context) map[string]interface{} {
	params := map[string]interface{}{}
	ctx.ReadJSON(&params)
	return params
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

func (c *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	//result["code"] = utils.RCODE_OK
	//result["message"] = utils.RecodeText(utils.RCODE_OK)

	if args != nil {
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

	c.Ctx.JSON(result)
	c.Ctx.StopExecution()
	return
}

// 解析分页数据
func (c *Common) ParsePageData(ctx iris.Context) (pageNum, pageSize int) {
	pageNum = cast.ToInt(ctx.URLParam("pagenum"))
	pageSize = cast.ToInt(ctx.URLParam("pagesize"))
	if pageNum == 0 {
		pageNum = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	if pageSize >= 1000 {
		pageSize = 1000
	}
	return pageNum, pageSize
}
