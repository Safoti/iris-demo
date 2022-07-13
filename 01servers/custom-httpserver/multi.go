package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description    定义多个服务器地址
 **/
func main() {
	 app:=iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})
	// Note: It's not needed if the first action is "go app.Run".
	if err := app.Build(); err != nil {
		panic(err)
	}
	// http://localhost:9090/
	// http://localhost:9090/mypath
	srv1 :=&http.Server{Addr: "：9090",Handler: app}
	go srv1.ListenAndServe()

	println("Start a server listening on http://localhost:9090")
	// http://localhost:9091/
	// http://localhost:9091/mypath
	srv2 := &http.Server{Addr: ":9091", Handler: app}
	go srv2.ListenAndServe()
	println("Start a server listening on http://localhost:9091")
	app.Listen(":8080") // Block here.
}
