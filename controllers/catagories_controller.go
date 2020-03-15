package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"vue_shop/common"
	"vue_shop/services"
)

type CategoriesController struct {
	Ctx     iris.Context
	Service *services.CategoriesService
	common.Common
}

func NewCategoriesController() *CategoriesController {
	return &CategoriesController{Service: services.NewCategoriesService()}
}

// 商品分类数据列表
func (c *CategoriesController) Get() {
	_type := cast.ToInt(c.Ctx.URLParam("type"))
	pageNum, pageSize := ParsePageData(c.Ctx)
	data := c.Service.List(_type, pageNum, pageSize)
	c.ReturnSuccess(data)
}

// 添加分类
func (c *CategoriesController) Post() {
	params := c.ParseParams(c.Ctx)
	if data, err:=c.Service.Create(params);err !=nil{
		c.ReturnJson(400, err.Error())
		return
	}else{
		c.ReturnSuccess(data)
	}

}

// 根据 id 查询分类
func (c *CategoriesController) GetBy(id int) {
	if data, err := c.Service.GetOne(id); err != nil {
		c.ReturnJson(400, err.Error())
	} else {
		c.ReturnSuccess(data)
	}
}

//  编辑提交分类
func (c *CategoriesController) PutBy(id int) {
	params := c.ParseParams(c.Ctx)
	if data, err := c.Service.Update(id, params); err != nil {
		c.ReturnJson(400, err.Error())
	} else {
		c.ReturnSuccess(data)
	}
}

// 删除分类
func (c *CategoriesController) DeleteBy(id int) {
	if err := c.Service.DeleteOne(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	}
	c.ReturnSuccess()
}

// 分类参数管理
func (c *CategoriesController) GetAttributes() {
	cateId, _ := c.Ctx.Params().GetInt("id")
	c.Ctx.WriteString(cast.ToString(cateId))
}

func ParsePageData(c iris.Context) (pageNum, pageSize int) {
	pageNum = cast.ToInt(c.URLParam("pagenum"))
	pageSize = cast.ToInt(c.URLParam("pagesize"))
	if pageNum == 0 && pageSize == 0 {
		pageNum = 1
		pageSize = 5000
	}
	return pageNum, pageSize
}
