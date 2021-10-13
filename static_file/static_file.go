package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()
	//set route
	app.HttpServer.SetEnabledListDir(true)
	app.HttpServer.ServerFile("/static/*filepath", "/www/public")
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
