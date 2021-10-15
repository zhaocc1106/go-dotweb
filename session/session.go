package main

import (
	"encoding/gob"
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/session"
)

type UserInfo struct {
	Name      string
	SessionId string
}

func StartServer() error {
	//init DotApp
	app := dotweb.New()

	//init session
	app.HttpServer.SetEnabledSession(true)
	//runtime session mode
	app.HttpServer.SetSessionConfig(session.NewDefaultRuntimeConfig())
	//redis session mode
	//app.HttpServer.SetSessionConfig(session.NewDefaultRedisConfig("redis://172.18.116.216:6379/0"))

	//set route
	app.HttpServer.GET("/", func(ctx dotweb.Context) error {
		fmt.Println("sessionId: " + ctx.SessionID())
		userSession := ctx.Session().Get(ctx.SessionID())
		if userSession != nil {
			userInfo := userSession.(UserInfo)
			fmt.Println("Have session: " + userInfo.Name + ", " + userInfo.SessionId)
			ctx.WriteString("Have session: " + userInfo.Name + ", " + userInfo.SessionId)
		} else {
			userSession = UserInfo{
				Name:      "zhaochaochao",
				SessionId: "123456",
			}
			err := ctx.Session().Set(ctx.SessionID(), userSession)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Not have session, set session.")
			ctx.WriteString("Not have session, set session.")
		}
		return nil
	})

	//begin server
	fmt.Println("Start http server.")
	err := app.StartServer(8080)
	return err
}

func init() {
	gob.Register(UserInfo{})
}

func main() {
	err := StartServer()
	if err != nil {
		return
	}
}
