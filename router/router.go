package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/accesslog"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()

	// 开启日志，设置日志目录
	app.SetEnabledLog(true)
	app.SetLogPath("/tmp/dotweb_server/")

	// 静态路由
	app.HttpServer.GET("/hello", func(context dotweb.Context) error {
		return context.WriteString("hello world!")
	})

	// 参数路由
	app.HttpServer.GET("/news/:category/:newsid", func(context dotweb.Context) error {
		category := context.GetRouterName("category")
		newsid := context.GetRouterName("newsid")
		return context.WriteString("news info: category: " + category + ", newsid: " + newsid)
	})

	// 组路由
	group := app.HttpServer.Group("/admin")
	group.Use(accesslog.Middleware()) // 该组路由启用access log中间件
	group.GET("/index", func(context dotweb.Context) error {
		return context.WriteString("/admin/index")
	})

	// begin server
	fmt.Println("Start http server.")
	err := app.StartServer(8080)
	return err
}

func main() {
	err := StartServer()
	if err != nil {
		return
	}
}
