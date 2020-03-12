package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type RightsController struct {
	Ctx     iris.Context
	Service *services.RightsServices
	common.Common
}

func NewRightsController() *RightsController {
	return &RightsController{Service: services.NewRightsServices()}
}

func (this *RightsController) GetBy(_type string) {
	if data, err := this.Service.RightList(_type); err != nil {
		this.ReturnJson(500, err.Error())
		return
	} else {
		this.ReturnSuccess(data)
	}
}
