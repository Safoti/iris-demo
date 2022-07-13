package main

import (
	"github.com/kataras/iris/v12"
	"net"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description   自定义监听
 **/
func main() {
	app := iris.New()
	app.Get("/", func(context iris.Context) {
		context.Writef("hello from the server")
	})
	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s",	 ctx.Path())
	})
    //自定义监听
	l, err := net.Listen("tcp4", ":8099")
	if err != nil {
		panic(err)
	}

	app.Run(iris.Listener(l))

}
