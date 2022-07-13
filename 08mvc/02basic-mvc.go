package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/
func main() {
    app:=iris.New()
    //设置日志级别
	app.Logger().SetLevel("debug")
	basic:=app.Party("/basic")
	{
		mvc.Configure(basic, basicMVC)
	}
	app.Listen(":8080")
}
func basicMVC(app *mvc.Application){
	// GET: http://localhost:8080/basic
	// GET: http://localhost:8080/basic/custom
	// GET: http://localhost:8080/basic/custom2
	app.Handle(new(basicController))

	// GET: http://localhost:8080/basic/sub
	app.Party("/sub").Handle(new(basicSubController))

}
type LoggerService interface {
	Log(string)
}
type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger)log(msg string)  {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}

type basicController struct {
	Logger LoggerService
	Seesion *sessions.Session
}
func(c *basicController)beforeActivation(b mvc.BeforeActivation){
	b.HandleMany("GET", "/custom /custom2", "Custom")

}
func (c *basicController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}

func (c *basicController)Get()string  {
	count:=c.Seesion.Increment("count",1)
	body := fmt.Sprintf("Hello from basicController\nTotal visits from you: %d", count)
	c.Logger.Log(body)
	return body
}

func (c *basicController)Custom()string  {
	return "custom"
}

type basicSubController struct {
	Session sessions.Session
}

func (c *basicSubController)Get()string  {
	count := c.Session.GetIntDefault("count", 1)
	return fmt.Sprintf("Hello from basicSubController.\nRead-only visits count: %d", count)
}