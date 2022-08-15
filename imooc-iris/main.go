package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"imooc/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 注册视图
	app.RegisterView(iris.HTML("./web/views", ".html"))
	// 注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	app.Run(iris.Addr(":4000"))
}