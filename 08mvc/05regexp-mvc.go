package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description   使用正则表达式
 **/

func main() {
	app := iris.New()

	mvcApp := mvc.New(app.Party("/module"))
	mvcApp.Handle(new(myController))

	// http://localhost:8080/module/xxx.json (OK)
	// http://localhost:8080/module/xxx.xml  (Not Found)
	app.Listen(":8080")
}

type myController  struct {
	
}

func (m *myController)BeforeActive(b mvc.BeforeActivation)  {
	//下面是英文解释
	// b.Dependencies().Register
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the HandleJSON

	// "^[a-zA-Z0-9_.-]+.json$)" to validate file-name pattern and json
	// or just:  ".json$" to validate suffix.

	b.Handle("GET","/{file:string regexp(^[a-zA-Z0-9_.-]+.json$))}","HandleJSON")
}
func (m *myController)HandleJSON(file string) string {
	return "custom serving of json: " + file
}