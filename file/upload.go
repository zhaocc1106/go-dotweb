package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
)

func InitRoute(server *dotweb.HttpServer) {
	// curl -F 'file=@file_name' 127.0.0.1:8080/file
	server.Router().POST("/file", FileUpload)
}

func FileUpload(ctx dotweb.Context) error {
	upload, err := ctx.Request().FormFile("file")
	if err != nil {
		return ctx.WriteString("FormFile error " + err.Error())
	} else {
		_, err = upload.SaveFile("/tmp/" + upload.FileName())
		if err != nil {
			return ctx.WriteString("SaveFile error => " + err.Error())
		} else {
			return ctx.WriteString("SaveFile success || " + upload.FileName() + " || " + upload.GetFileExt() + " || " + fmt.Sprint(upload.Size()))
		}
	}
}

func main() {
	//初始化DotServer
	app := dotweb.New()

	//启用开发模式
	app.SetDevelopmentMode()
	//启用访问日志
	app.SetEnabledLog(true)
	app.UseRequestLog()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}
