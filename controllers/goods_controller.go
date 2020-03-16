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
