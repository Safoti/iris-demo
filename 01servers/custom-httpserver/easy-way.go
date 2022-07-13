package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description  简单方式
 **/
func main() {
  app:=iris.New()
  app.Get("/", func(ctx iris.Context) {
	 ctx.Writef("hello from the server")
  })
  app.Get("/mypath", func(ctx iris.Context) {
	  ctx.Writef("Hello from %s", ctx.Path())
  })
	// Any custom fields here. Handler and ErrorLog are set to the server automatically

	srv := &http.Server{Addr: ":8090"}
	app.Run(iris.Server(srv)) // // same as app.Listen(":8080")


}
