package main

import "github.com/kataras/iris/v12"

/**
 * @Author safoti
 * @Date Created in 2022/7/12
 * @Description
 **/
func main() {
	app:=iris.New()
	app.OnErrorCode(iris.StatusNotFound,notFound)
	app.OnErrorCode(iris.StatusInternalServerError,internalServerError)
	// 为所有的大于等于400的状态码注册一个处理器：
	// app.OnAnyErrorCode(handler)
	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.View("errors/404.html")
}
func index(ctx iris.Context) {
	ctx.View("index.html")
}