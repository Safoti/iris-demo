package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {

	app := iris.Default()
	//获取请求路径中的参数
	app.Get("/welcome", func(ctx context.Context) {
		first := ctx.URLParamDefault("firstname", "lktbz")
		last := ctx.URLParam("lastname")
		ctx.Writef("Hello %s %s",first,last)
	})
	app.Listen(":8080")
}
