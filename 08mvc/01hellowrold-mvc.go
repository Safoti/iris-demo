package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/mvc"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/12
 * @Description   iris mvc  helloworld
 **/
func main() {
 	app:=newApp()
 	app.Listen(":9090")
}
// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct {}

func newApp() *iris.Application {
	app := iris.New()

	app.Use(logger.New())
	// Serve a controller based on the root Router, "/".
	//注册controller
	mvc.New(app).Handle(new(ExampleController))
	return  app
}

//service
// Method:   GET
// Resource: http://localhost:9090
//响应方式html
func (c *ExampleController)Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text: "<h1>Welcome</h1>",
	}
}
// GetPing serves
// Method:   GET
// Resource: http://localhost:9090/ping
//响应方式 string
func (c *ExampleController)GetPing()string  {
	return "pong"
}
// GetHello serves
// Method:   GET
// Resource: http://localhost:9090/hello
//响应方式 map
func (c *ExampleController)GetHello()interface{}  {
	return map[string]string{"message":"Hello Iris!"}
}



