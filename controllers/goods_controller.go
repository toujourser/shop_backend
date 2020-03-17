package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type GoodsController struct {
	Service *services.GoodsServices
	Ctx     iris.Context
	common.Common
}

func NewGoodsController() *GoodsController {
	return &GoodsController{Service: services.NewGoodsServices()}
}

// 商品列表数据
func (c *GoodsController) Get() {
	query := c.Ctx.URLParam("query")
	pageNum, pageSize := c.ParsePageData(c.Ctx)

	if data, err := c.Service.List(pageNum, pageSize, query); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 添加商品
func (c *GoodsController) Post() {

	if data, err := c.Service.Create(c.Ctx); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 根据 ID 查询商品
func (c *GoodsController) GetBy(id int) {
	if data, err := c.Service.GetOne(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 编辑提交商品
func (c *GoodsController) PutBy(id int) {
	if data, err := c.Service.Update(id, c.Ctx); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 删除商品
func (c *GoodsController) DeleteBy(id int) {
	if err := c.Service.Delete(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess()
	}
}
