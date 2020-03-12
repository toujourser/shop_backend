package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"vue_shop/controllers"
)

func InitRoute(app *iris.Application) {
	mvc.New(app.Party("/api/v1.0/user")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/api/v1.0/user/{uid:int}")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/api/v1.0/user/{uid:int}/state/{state:boolean}")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/api/v1.0/menus")).Handle(controllers.NewMenusController())
	mvc.New(app.Party("/api/v1.0/rights")).Handle(controllers.NewRightsController())
	mvc.New(app.Party("/api/v1.0/roles")).Handle(controllers.NewRolesController())
	mvc.New(app.Party("/api/v1.0/roles/{roleId:int}/rights/{rightId:int}")).Handle(controllers.NewRolesController())
	mvc.New(app.Party("/api/v1.0/roles/{roleId:int}")).Handle(controllers.NewRolesController())
}
