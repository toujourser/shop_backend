package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"vue_shop/config"
	"vue_shop/middleware"
	models "vue_shop/models"
	"vue_shop/route"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		fmt.Println("Err: ", err)
		panic(err)
	}
	//go services.SyncDashBoard()
	models.DB.Init()

	app := newApp()

	route.InitRoute(app)

	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	// 404 错误处理
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>404 Resource Not found</b>")
	})
	// 500 错误处理
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})

	// 跨域访问
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	app.Use(crs)
	//app.Use(middleware.Cors)
	app.Use(middleware.GetJWT().Serve) // JWT
	app.AllowMethods(iris.MethodOptions)
	app.Configure(iris.WithOptimizations, iris.WithoutInterruptHandler)
	return app
}
