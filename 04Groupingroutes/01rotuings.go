package main
//
//import (
//	"github.com/kataras/iris/v12"
//	 "github.com/kataras/iris/v12/"
//)
//func main() {
//    app:=iris.Default()
//    //创建dev 组
//  	dev:=	app.Party("/dev")
//	//创建test 组
//  	test:=app.Party("/test")
//	//创建请求方法
//  	dev.Get("/login",loginDemo)
//	//创建请求方法
//  	test.Get("/login",loginDemo)
//	app.Listen(":8080")
//}
//func loginDemo(ctx .Context) {
//		cas:=	ctx.Request().Method
//		ctx.Writef("分组测试的请求方位是？：%s",cas)
//}
