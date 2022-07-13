package main

import "github.com/kataras/iris/v12"

type MyType struct {
	Name string `url:"name,required"`
	Age  int    `url:"age"`
}


func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		var t MyType
		err:=ctx.ReadQuery(&t)
		if err!=nil{
			panic(err)
			return
		}
		ctx.Writef("MyType: %#v", t)
	})

	app.Listen(":8080")
}
