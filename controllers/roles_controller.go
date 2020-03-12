package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type RolesController struct {
	Ctx     iris.Context
	Service *services.RolesService
	common.Common
}

func NewRolesController() *RolesController {
	return &RolesController{Service: services.NewRolesService()}
}

// 角色列表
func (c *RolesController) Get() {
	if data, err := c.Service.List(); err != nil {
		c.ReturnJson(500, err.Error())
	} else {
		c.ReturnSuccess(data)
	}
}

// 添加角色
func (c *RolesController) Post() {
	args := map[string]string{}
	if err := c.Ctx.ReadJSON(&args); err != nil {
		c.ReturnJson(400, "参数错误")
		return
	}
	if data, err := c.Service.Create(args["roleName"], args["roleDesc"]); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 根据 ID 查询角色
func (c *RolesController) GetBy(id int) {
	if data, err := c.Service.GetOne(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 编辑角色
func (c *RolesController) PutBy(id int) {
	args := map[string]string{}
	if err := c.Ctx.ReadJSON(&args); err != nil {
		c.ReturnJson(400, "参数错误")
		return
	}
	if data, err := c.Service.Update(id, args["roleName"], args["roleDesc"]); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 删除角色
func (c *RolesController) DeleteBy(id int) {
	if err := c.Service.DeleteOne(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	}
	c.ReturnSuccess()
}

// 删除角色指定权限
func (c *RolesController) Delete() {
	roleId, _ := c.Ctx.Params().GetInt("roleId")
	rightId, _ := c.Ctx.Params().GetInt("rightId")

	if data, err := c.Service.DeleteRight(roleId, rightId); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(data)
	}
}

// 角色授权
func (c *RolesController) PostRights() {
	roleId, _ := c.Ctx.Params().GetInt("roleId")
	args := map[string]string{}
	if err := c.Ctx.ReadJSON(&args); err != nil {
		c.ReturnJson(400, "参数错误")
		return
	}
	if err := c.Service.RoleImpower(roleId, args["rids"]); err != nil {
		c.ReturnJson(400, err.Error())
		return
	}
	c.ReturnSuccess()
}
