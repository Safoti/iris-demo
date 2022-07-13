package main

import "github.com/kataras/iris/v12"

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}


func main() {
	app := iris.New()
	// set the view html template engine
	app.RegisterView(iris.HTML("02 requests/templates", ".html").Reload(true))
	//路径跳转
	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			panic(err)
			return
		}
	})
	//获取表单数据
	app.Post("/form_action", func(ctx iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			if !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
				panic(err)
				return
			}
		}

		ctx.Writef("Visitor: %#v", visitor)
	})
	app.Listen(":8080")

}
