package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.Default()
	//获取表单数据
	app.Post("/form_post", func(ctx context.Context) {
		message :=	ctx.PostValue("message")
		nick := ctx.PostValueDefault("nick", "anonymous")
	    ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	app.Listen(":8080")
}
