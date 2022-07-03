package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

/**
   表单提交获取请求参数，并获取表单数据
 */
func main() {
	app := iris.Default()
	app.Post("/post", func(ctx context.Context) {
		id, err := ctx.URLParamInt("id")
		if err != nil {
			return
		}
		page := ctx.URLParamIntDefault("page", 0)
		name := ctx.PostValue("name")
		message := ctx.PostValue("message")
		ctx.Writef("id: %d; page: %d; name: %s; message: %s", id, page, name, message)
	})
	app.Listen(":8080")

}
