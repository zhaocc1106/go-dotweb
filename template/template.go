package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func StartServer() error {
	app := dotweb.New()
	app.HttpServer.GET("/", func(ctx dotweb.Context) error {
		fmt.Println("Get /")
		ctx.ViewData().Set("data", "图书信息")

		type BookInfo struct {
			Name string
			Size int64
		}

		m := make([]*BookInfo, 5)
		m[0] = &BookInfo{Name: "book0", Size: 1}
		m[1] = &BookInfo{Name: "book1", Size: 10}
		m[2] = &BookInfo{Name: "book2", Size: 100}
		m[3] = &BookInfo{Name: "book3", Size: 1000}
		m[4] = &BookInfo{Name: "book4", Size: 10000}
		ctx.ViewData().Set("Books", m)

		err := ctx.View("template/testview.html")
		return err
	})
	return app.StartServer(8080)
}

func main() {
	err := StartServer()
	if err != nil {
		return
	}
}
