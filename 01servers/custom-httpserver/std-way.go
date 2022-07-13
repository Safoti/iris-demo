package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/
func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})
	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})
	// call .Build before use the 'app' as a http.Handler on a custom http.Server
	app.Build()
	// create our custom server and assign the Handler/Router
	src:=&http.Server{Handler: app,Addr: ":8090"}
	println("Start a server listening on http://localhost:8080")
	src.ListenAndServe() // same as app.Listen(":8080")
}
