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
	pageNum, pageSize := c.ParsePageData(c.Ctx)
	data := c.Service.List(_type, pageNum, pageSize)
	c.ReturnSuccess(data)
}

// 添加分类
func (c *CategoriesController) Post() {
	params := c.ParseParams(c.Ctx)
	if data, err := c.Service.Create(params); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
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
	sel := c.Ctx.URLParam("sel")
	if data, err := c.Service.GetAttributes(cast.ToInt64(cateId), sel); err != nil {
		c.ReturnJson(400, err.Error())
	} else {
		c.ReturnSuccess(data)
	}

}

// 添加动态参数或者静态属性
func (c *CategoriesController) PostAttributes() {
	cateId, _ := c.Ctx.Params().GetInt("id")
	params := c.ParseParams(c.Ctx)

	if data, err := c.Service.CreateAttributes(cast.ToInt64(cateId), params); err != nil {
		c.ReturnJson(400, err.Error())
	} else {
		c.ReturnSuccess(data)
	}
}

// 删除分类属性参数
// url: /api/v1.0/categories/:id/attributes/:attrid
func (c *CategoriesController) DeleteAttributesBy(attrid int) {
	cateId, _ := c.Ctx.Params().GetInt("id")

	if err := c.Service.DeleteAttributes(cateId, attrid); err != nil {
		c.ReturnJson(400, err.Error())
		return
	}
	c.ReturnSuccess()
}

// 根据 ID 查询参数
func (c *CategoriesController) GetAttributesBy(attrId int) {
	cateId, _ := c.Ctx.Params().GetInt("id")
	sel := c.Ctx.URLParam("attr_sel")
	vals := c.Ctx.URLParam("attr_vals")
	if data, err := c.Service.GetAttributesByAttrId(cateId, attrId, sel, vals); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 编辑提交参数
func (c *CategoriesController) PutAttributesBy(attrId int) {
	cateId, _ := c.Ctx.Params().GetInt("id")
	params := c.ParseParams(c.Ctx)
	if data, err := c.Service.PutAttributesByAttrId(cateId, attrId, params); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}
