package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type MenusController struct {
	Ctx iris.Context
	common.Common
	Service *services.MenusServices
}

func NewMenusController() *MenusController {
	return &MenusController{Service: services.NewMenusServices()}
}

func (c *MenusController) Get() {
	menus, err := c.Service.List()
	if err != nil {
		c.Common.ReturnJson(500, err.Error())
		return
	}
	c.Common.ReturnSuccess(menus)
}
