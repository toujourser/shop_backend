package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"vue_shop/controllers"
)

func InitRoute(app *iris.Application) {
	mvc.New(app.Party("/api/v1.0/user")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/api/v1.0/user/{uid:int}/state/{state:boolean}")).Handle(controllers.NewUserController())
	mvc.New(app.Party("/api/v1.0/menus")).Handle(controllers.NewMenusController())
}
