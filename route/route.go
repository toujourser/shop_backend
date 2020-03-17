package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"vue_shop/controllers"
)

func InitRoute(app *iris.Application) {
	api := mvc.New(app.Party("/api/v1.0"))
	api.Party("/user").Handle(controllers.NewUserController())
	api.Party("/user/{uid:int}").Handle(controllers.NewUserController())
	api.Party("/user/{uid:int}/state/{state:boolean}").Handle(controllers.NewUserController())
	api.Party("/menus").Handle(controllers.NewMenusController())
	api.Party("/rights").Handle(controllers.NewRightsController())
	api.Party("/role").Handle(controllers.NewRolesController()) // 根据ID获取单个角色 删除角色
	api.Party("/roles").Handle(controllers.NewRolesController())
	api.Party("/roles/{roleId:int}").Handle(controllers.NewRolesController())

	api.Party("/categories").Handle(controllers.NewCategoriesController())
	api.Party("/categorie/{id:int}").Handle(controllers.NewCategoriesController()) // 分类参数管理

	api.Party("/goods").Handle(controllers.NewGoodsController())
	api.Party("/reports/type/1").Handle(controllers.NewReportsController())
}
