package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()

	//set route
	app.HttpServer.GET("/get", func(ctx dotweb.Context) error{
		// curl 127.0.0.1/get?name=test
		name := ctx.QueryString("name")
		return ctx.WriteString("name: " + name)
	})
	app.HttpServer.POST("/post", func(ctx dotweb.Context) error {
		// curl -X POST -F "name=test" 127.0.0.1:8080/post
		name := ctx.PostFormValue("name")
		//body := ctx.Request().PostBody()
		//fmt.Println("body: " + string(body))
		return ctx.WriteString("name: " + name)
	})
	app.HttpServer.POST("/bind", func(ctx dotweb.Context) error {
		// curl -X POST -H 'Content-Type: application/json' -d '{"UserName":"test", "Sex":1}' 127.0.0.1:8080/bind
		type UserInfo struct {
			UserName string `json:"name"`
			Sex      int `json:"sex"`
		}
		user := new(UserInfo)
		if err := ctx.Bind(user); err != nil {
			fmt.Println(err)
			return err
		}
		return ctx.WriteJson(user)
	})

	//begin server
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
