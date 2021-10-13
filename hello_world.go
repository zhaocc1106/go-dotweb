package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()
	//set route
	app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
		return ctx.WriteString("welcome to my first web!")
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
