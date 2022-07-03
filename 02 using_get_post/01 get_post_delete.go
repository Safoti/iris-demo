package main

import "github.com/kataras/iris/v12"

func main() {
	//创建请求方式
	app := iris.Default()
	app.Get("/someGet")
	app.Post("/someGet")
	app.Put("/someGet")
    app.Listen(":8080")

}
