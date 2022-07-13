package main

import "github.com/kataras/iris/v12"

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/
func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	app.Listen(":8080", iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "utf-8",
	}))
}
