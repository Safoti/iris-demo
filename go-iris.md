# 								go-iris 

> 刚开始必须helloworld

# 一：hello iris

```go
package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	// 输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080/welcome
	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	// 输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	// 输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Run(iris.Addr(":8080")) //8080 监听端口
}

```

**本手册学习过程为：**

接收数据

- 请求路径以及携带的参数

响应数据

- 后台处理的数据以合适的方式响应回去

# 二：接收数据

### 2.1：请求类型



```go
package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)
func main() {
/**
	  iris 的请求方式  Get Post Put Delete
	  http://localhost:8080/get
 */
	app := iris.Default()
	app.Get("/get", func(ctx context.Context) {
		ctx.WriteString("get 浏览器响应信息  ")
	})
	app.Post("/post", func(ctx context.Context) {
		ctx.WriteString("post 浏览器响应信息")
	})
	app.Put("/put", func(ctx context.Context) {
		ctx.WriteString("put  浏览器响应信息")
	})
	app.Delete("/delete", func(ctx context.Context) {
		ctx.WriteString("delete 浏览器响应信息")
	})
	app.Patch("/patch", func(ctx context.Context) {
		ctx.WriteString("patch 浏览器响应信息")
	})
	app.Options("/options", func(ctx context.Context) {
		ctx.WriteString("options 浏览器响应信息")
	})
//监听端口
app.Listen(":8080")

}
```

### 2.2：路径参数

```go
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
```

### 2.3：绑定 json 

```go

type Company struct {
	Name  string
	City  string
	Other string
}
func MyHandler(ctx iris.Context) {
	var c Company

	if err := ctx.ReadJSON(&c); err != nil {
		panic(err)
		return
	}
	ctx.Writef("Received: %#+v\n", c)
}

func main() {
	app := iris.New()
	app.Post("/", MyHandler)
	app.Listen(":8080", iris.WithOptimizations)
}

```

读取json 流

html

```html
<!DOCTYPE html>
<head>
    <meta charset="utf-8">
</head>
<body>
<form action="/form_action" method="post">
    Username: <input type="text" name="Username" /> <br />
    Mail: <input type="text" name="Mail" /> <br />
    Select one or more:  <br/>
    <select multiple="multiple" name="mydata">
        <option value='one'>One</option>
        <option value='two'>Two</option>
        <option value='three'>Three</option>
        <option value='four'>Four</option>
    </select>

    <hr />
    <input type="submit" value="Send data" />

</form>
</body>
</html>
```

go

```go
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
    //以及设置html 存放路径
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

```

### 2.8：路由

> 相同请求，不同的分组进行隔离

```GO
import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)
func main() {
    app:=iris.Default()
    //创建dev 组
  	dev:=	app.Party("/dev")
	//创建test 组
  	test:=app.Party("/test")
	//创建请求方法
  	dev.Get("/login",loginDemo)
	//创建请求方法
  	test.Get("/login",loginDemo)
	app.Listen(":8080")
}
func loginDemo(ctx context.Context) {
		cas:=	ctx.Request().Method
		ctx.Writef("分组测试的请求方位是？：%s",cas)
}

```

# 三：响应数据

### 3.1：内容协商

> 可以缓缓，后面分析其用处

```go
package main

import "github.com/kataras/iris/v12"

//内容协商、
type testdata struct {
	Name string `json:"name" xml:"Name"`
	Age  int    `json:"age" xml:"Age"`
}

func main() {
	app := newApp()
	app.Listen(":8080")
}

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/resource", func(ctx iris.Context) {
		data := testdata{
			Name: "test name",
			Age:  26,
		}

		// Server allows response only JSON and XML. These values
		// are compared with the clients mime needs. Iris comes with default mime types responses
		// but you can add a custom one by the `Negotiation().Mime(mime, content)` method,
		// same for the "accept".
		// You can also pass a custom ContentSelector(mime string) or ContentNegotiator to the
		// `Context.Negotiate` method if you want to perform more advanced things.
		//
		//
		// By-default the client accept mime is retrieved by the "Accept" header
		// Indeed you can override or update it by `Negotiation().Accept.XXX` i.e
		// ctx.Negotiation().Accept.Override().XML()
		//
		// All these values can change inside middlewares, the `Negotiation().Override()` and `.Accept.Override()`
		// can override any previously set values.
		// Order matters, if the client accepts anything (*/*)
		// then the first prioritized mime's response data will be rendered.
		ctx.Negotiation().JSON().XML()
		// Accept-Charset vs:
		ctx.Negotiation().Charset("utf-8", "iso-8859-7")

		// Alternatively you can define the content/data per mime type
		// anywhere in the handlers chain using the optional "v" variadic
		// input argument of the Context.Negotiation().JSON,XML,YAML,Binary,Text,HTML(...) and e.t.c
		// example (order matters):
		// ctx.Negotiation().JSON(data).XML(data).Any("content for */*")
		// ctx.Negotiate(nil)

		// if not nil passed in the `Context.Negotiate` method
		// then it overrides any contents made by the negotitation builder above.
		_, err := ctx.Negotiate(data)
		if err !=nil{
			ctx.Writef("%v", err)
		}
	})


	app.Get("/resource2", func(ctx iris.Context) {
		jsonAndXML := testdata{
			Name: "test name",
			Age:  26,
		}
		// I prefer that one, as it gives me the freedom to modify
		// response data per accepted mime content type on middlewares as well.
		ctx.Negotiation().JSON(jsonAndXML).XML(jsonAndXML).HTML("<h1>Test Name</h1><h2>Age 26</h2>")
		ctx.Negotiate(nil)
	})

	app.Get("/resource3", func(ctx iris.Context) {

		// If that line is missing and the requested
		// mime type of content is */* or application/xml or application/json
		// then 406 Not Acceptable http error code will be rendered instead.
		//
		// We also add the "gzip" algorithm as an option to encode
		// resources on send.
		ctx.Negotiation().JSON().XML().HTML().EncodingGzip()
		jsonAndXML := testdata{
			Name: "test name",
			Age:  26,
		}
		ctx.Negotiate(iris.N{
			// Text: for text/plain,
			// Markdown: for text/mardown,
			// Binary: for application/octet-stream,
			// YAML: for application/x-yaml,
			// JSONP: for text/javascript
			// Other: for anything else,
			JSON: jsonAndXML,                          // for application/json
			XML:  jsonAndXML,                          // for application/xml or text/xml
			HTML: "<h1>Test Name</h1><h2>Age 26</h2>", // for text/html
		})
	})
	return app
}

```

### 3.2：响应的数据类型

```go
package main

import (
	"encoding/xml"
	"fmt"
	"github.com/kataras/iris/v12"
)

// User example struct for json and msgpack.
type User struct {
	Firstname string `json:"firstname" msgpack:"firstname"`
	Lastname  string `json:"lastname" msgpack:"lastname"`
	City      string `json:"city" msgpack:"city"`
	Age       int    `json:"age" msgpack:"age"`
}

// ExampleXML just a test struct to view represents xml content-type
type ExampleXML struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

// ExampleYAML just a test struct to write yaml to the client.
type ExampleYAML struct {
	Name       string `yaml:"name"`
	ServerAddr string `yaml:"ServerAddr"`
}

func main() {
	app := iris.New()
	// Read
	app.Post("/decode", func(ctx iris.Context) {
		var usr User
		ctx.ReadJSON(&usr)
		fmt.Println(usr)
		ctx.Writef("%s %s is %d years old and comes from %s!", usr.Firstname, usr.Lastname, usr.Age, usr.City)

	})

	// Write
	app.Post("/encode", func(ctx iris.Context) {
		u := User{
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}
		// Manually setting a content type: ctx.ContentType("text/javascript")
		ctx.JSON(u)
	})
	// Use Secure field to prevent json hijacking.
	// It prepends `"while(1),"` to the body when the data is array.
    app.Get("/json_secure", func(ctx iris.Context) {
		respo:= []string{"val1", "val2", "val3"}
		options:=iris.JSON{Indent:""}
		ctx.JSON(respo,options)
	})
	// Use ASCII field to generate ASCII-only JSON
	// with escaped non-ASCII characters.
	app.Get("/json_ascii", func(ctx iris.Context) {
		response := iris.Map{"lang": "GO-虹膜", "tag": "<br>"}
		options := iris.JSON{Indent: "    "}
		ctx.JSON(response, options)

		/* Will output:
		   {
		       "lang": "GO-\u8679\u819c",
		       "tag": "\u003cbr\u003e"
		   }
		*/
	})

	app.Get("/json_raw", func(ctx iris.Context) {
		options := iris.JSON{UnescapeHTML: true}
		ctx.JSON(iris.Map{
			"html": "<b>Hello, world!</b>",
		}, options)

		// Will output: {"html":"<b>Hello, world!</b>"}
	})
	app.Get("/binary", func(ctx iris.Context) {
		// useful when you want force-download of contents of raw bytes form.
		ctx.Binary([]byte("Some binary data here."))
	})

	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Plain text here")
	})

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
	})

	//app.Get("/jsonp", func(ctx iris.Context) {
	//	ctx.JSONP(map[string]string{"hello": "jsonp"}, iris.JSONP{Callback: "callbackName"})
	//})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(ExampleXML{One: "hello", Two: "xml"})
		// OR:
		// ctx.XML(iris.XMLMap("keys", iris.Map{"key": "value"}))
	})

	app.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
	})

	app.Get("/yaml", func(ctx iris.Context) {
		ctx.YAML(ExampleYAML{Name: "Iris", ServerAddr: "localhost:8080"})
		// OR:
		// ctx.YAML(iris.Map{"name": "Iris", "serverAddr": "localhost:8080"})
	})

	// app.Get("/protobuf", func(ctx iris.Context) {
	// 	ctx.Protobuf(proto.Message)
	// })

	//app.Get("/msgpack", func(ctx iris.Context) {
	//	u := User{
	//		Firstname: "John",
	//		Lastname:  "Doe",
	//		City:      "Neither FBI knows!!!",
	//		Age:       25,
	//	}
	//
	//	ctx.MsgPack(u)
	//})
	// Other content types,


	app.Listen(":8080", iris.WithOptimizations)
}

```



接收数据

- 请求过来所携带的数据

响应数据

