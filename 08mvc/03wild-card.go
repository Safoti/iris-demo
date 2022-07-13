package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description   通配符写法
 **/
func main() {
   app:=iris.New()
   userRou:=app.Party("/users")
   //处理某个请求
   mvc.New(userRou).Handle(new(MyController))
   /**
     等价于
    usersRouter.Get("/{p:path}", func(ctx iris.Context) {
   	// 	wildcardPathParameter := ctx.Params().Get("p")
   	// 	ctx.JSON(response{
   	// 		Message: "The path parameter is: " + wildcardPathParameter,
   	// 	})
   	// })
   	/*
   		curl --location --request GET 'http://localhost:8080/users/path_segment_1/path_segment_2'
   		Expected Output:
   		{
   		  "message": "The wildcard is: path_segment_1/path_segment_2"
   		}


    */
 app.Listen(":9090")
}
//定义controller
type MyController struct {}
//自定义响应信息
type response struct {
	Message string `json:"message"`
}
//service
func (c *MyController)GetByWildcard(wp string)response{
	return response{
		Message: "The path parameter is: " + wp,
	}
}