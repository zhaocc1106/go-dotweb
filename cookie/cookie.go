package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"net/http"
)

func StartServer() error {
	//init DotApp
	app := dotweb.New()
	//set route
	app.HttpServer.GET("/", func(ctx dotweb.Context) error {
		sessionId, err := ctx.ReadCookie("session_id")
		if err != nil {
			fmt.Println(err)
			ctx.SetCookie(&http.Cookie{Name: "token",
				Value:    "this is token",
				MaxAge:   10,
				HttpOnly: true,
				Secure:   true,
			})
			ctx.SetCookie(&http.Cookie{Name: "session_id",
				Value:    "this is session id",
				MaxAge:   10,
				HttpOnly: true,
				Secure:   true,
			})
			return ctx.WriteString("No login, set token and session_id")
		}

		fmt.Println(sessionId)

		return ctx.WriteString("Have login with session_id: " + sessionId.Value)
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
