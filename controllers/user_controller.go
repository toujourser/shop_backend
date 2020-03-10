package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"vue_shop/common"
	"vue_shop/middleware"
	"vue_shop/services"
)

type UserController struct {
	Ctx     iris.Context
	Service *services.UserServices
	common.Common
}

func NewUserController() *UserController {
	return &UserController{Service: services.NewAuthServices()}
}

// 用户登陆
func (c *UserController) PostLogin() {
	jsonData := make(map[string]string)

	if err := c.Ctx.ReadJSON(&jsonData); err != nil {
		c.ReturnJson(400, cast.ToString(err))
		return
	}
	if user, err := c.Service.UserLogin(jsonData); err != nil {
		c.ReturnJson(401, cast.ToString(err))
		return
	} else {
		token := middleware.GenrateAdminToken(user)
		data := map[string]interface{}{
			"id":       user.MgId,
			"rid":      user.RoleId,
			"username": user.MgName,
			"mobile":   user.MgMobile,
			"email":    user.MgEmail,
			"token":    token,
		}
		c.ReturnJson(200, "登录成功", data)
		return
	}

}

// 获取符合条件的用户
func (c *UserController) GetList() {
	query := c.Ctx.URLParam("query")
	pageNum := cast.ToInt(c.Ctx.URLParam("pagenum"))
	pageSize := cast.ToInt(c.Ctx.URLParam("pagesize"))

	userList, total := c.Service.UserList(pageNum, pageSize, query)
	data := map[string]interface{}{
		"users":   userList,
		"total":   total,
		"pagenum": pageNum,
	}
	c.Common.ReturnSuccess(data)
}

// 根据用户ID查询用户信息
func (c *UserController) GetBy(id int) {
	if u, err := c.Service.GetOne(id); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(u)
	}
}

// 修改用户信息
func (c *UserController) PutBy(id int) {
	info := map[string]string{}
	c.Ctx.ReadJSON(&info)

	if u, err := c.Service.PutOne(id, info["email"], info["mobile"]); err != nil {
		c.ReturnJson(400, err.Error())
		return
	} else {
		c.ReturnSuccess(u)
	}

}

// 切换用户状态
func (c *UserController) Put() {
	userId, _ := c.Ctx.Params().GetInt("uid")
	userState, _ := c.Ctx.Params().GetBool("state")
	println("----", userId, userState)

	if data, err := c.Service.UserState(userId, userState); err != nil {
		c.Common.ReturnJson(500, err.Error())
		return
	} else {
		c.Common.ReturnSuccess(data)
	}

}

// 创建新用户
func (c *UserController) Post() {
	m := map[string]string{}
	if err := c.Ctx.ReadJSON(&m); err != nil {
		c.ReturnJson(400, cast.ToString(err))
		return
	}

	if m["username"] == "" || m["password"] == "" || m["email"] == "" || m["mobile"] == "" {
		c.ReturnJson(400, "请求参数错误")
		return
	}

	if data, err := c.Service.Create(m); err != nil {
		c.ReturnJson(500, err.Error())
		return
	} else {
		c.ReturnJson(201, "创建成功", data)
	}
}

func (c *UserController) DeleteBy(id int) {
	if err := c.Service.DeleteOne(id); err != nil {
		c.ReturnJson(500, err.Error())
		return
	}
	c.ReturnSuccess()
}
