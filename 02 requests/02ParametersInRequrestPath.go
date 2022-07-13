package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

/**
  请求路径参数
 */
func main() {
	app := iris.Default()
		//  匹配请求地址 http://localhost:8080/user/lktbz
	// This handler will match /user/john but will not match /user/ or /user
	app.Get("/user/{name}", func(ctx context.Context) {
		name:=ctx.Params().Get("name")
		fmt.Println(name)
		ctx.Writef("Hello %s", name)
	})
         //http://localhost:8080/user/lktbz/send 匹配请求地址
	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
     app.Get("/user/{name}/{action:path}", func(c context.Context) {
		 name := c.Params().Get("name")
		 action:=c.Params().Get("action")
		 message:=name+" :is:"+action
		 c.WriteString(message)
     })

	//参数带默认值的方式
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  http://localhost:8080/welcome?first=Jane&last=Doe
	app.Get("/welcome", func(ctx context.Context) {
	first:=	ctx.URLParamDefault("first","Guest")
	last:=ctx.URLParam("last")  //==	ctx.Request().URL.Query().Get("name")
		ctx.Writef("Hello %s %s", first, last)
	})

	/**
	   form 表单
	 */
	app.Post("/form", func(ctx context.Context) {
		message := ctx.PostValue("message")
		nick := ctx.PostValueDefault("nick","anonymous")
		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	app.Listen(":8080")

}
