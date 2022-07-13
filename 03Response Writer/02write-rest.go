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
