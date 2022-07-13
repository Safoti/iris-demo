package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
/**
	  iris 的请求方式
 */
	app := iris.Default()
	/**
		func Default() *Application {
			app := New()
			app.Use(recover.New())
			app.Use(requestLogger.New())
			app.defaultMode = true

			return app

	}
	 */
	app.Get("/get", func(ctx context.Context) {
		ctx.WriteString("get 浏览器响应信息  ")
	})
	app.Post("/post", func(ctx context.Context) {
		ctx.WriteString("post 浏览器响应信息")
	})
	app.Put("/put", func(ctx context.Context) {
		ctx.WriteString("put  浏览器响应信息")
	})

	app.Delete("/delete", func(ctx context.Context) {
		ctx.WriteString("delete 浏览器响应信息")
	})

	app.Patch("/patch", func(ctx context.Context) {
		ctx.WriteString("patch 浏览器响应信息")
	})

	app.Options("/options", func(ctx context.Context) {
		ctx.WriteString("options 浏览器响应信息")
	})

app.Listen(":8080")


}
