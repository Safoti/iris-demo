package main

import (
	"github.com/kataras/iris/v12"

)

type Company struct {
	Name  string
	City  string
	Other string
}
func MyHandler(ctx iris.Context) {
	var c Company

	if err := ctx.ReadJSON(&c); err != nil {
		panic(err)
		return
	}
	ctx.Writef("Received: %#+v\n", c)
}

func main() {
	app := iris.New()
	app.Post("/", MyHandler)
	app.Listen(":8080", iris.WithOptimizations)
}
