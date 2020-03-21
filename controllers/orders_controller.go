package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"vue_shop/common"
	"vue_shop/services"
)

type OrdersController struct {
	Ctx     iris.Context
	Service *services.OrdersServices
	common.Common
}

func NewOrdersController() *OrdersController {
	return &OrdersController{Service: services.NewOrdersServices()}
}
// 订单数据列表
func (c *OrdersController) Get() {
	query := c.Ctx.URLParam("query")
	pageNum, pageSize := c.ParsePageData(c.Ctx)
	user_id := c.Ctx.URLParam("user_id")
	pay_status := c.Ctx.URLParam("pay_status")
	is_send := c.Ctx.URLParam("is_send")
	order_fapiao_title := c.Ctx.URLParam("order_fapiao_title")
	order_fapiao_company := c.Ctx.URLParam("order_fapiao_company")
	order_fapiao_content := c.Ctx.URLParam("order_fapiao_content")
	consignee_addr := c.Ctx.URLParam("consignee_addr")

	if data, err := c.Service.List(pageNum, pageSize, cast.ToInt(user_id),
		query, pay_status, is_send, order_fapiao_title, order_fapiao_company,
		order_fapiao_content, consignee_addr); err != nil {
		c.ReturnJson(400, err.Error())
	} else {
		c.ReturnSuccess(data)
	}

}
