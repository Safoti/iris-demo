package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"sync/atomic"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/

func main() {
   app:=iris.New()
   mvc.New(app.Party("/")).Handle(&globalVisitorController{0})
	// http://localhost:8080
	app.Listen(":8080")
}

type globalVisitorController struct {
	visits uint64
}
func (c *globalVisitorController)Get()string{
	count := atomic.AddUint64(&c.visits, 1)
	return fmt.Sprintf("Total visitors: %d", count)
}