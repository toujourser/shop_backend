package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type ReportsController struct {
	Ctx     iris.Context
	Service *services.ReportsServices
	common.Common
}

func NewReportsController() *ReportsController {
	return &ReportsController{Service: services.NewReportsServices()}
}

func (c *ReportsController)Get(){

}
