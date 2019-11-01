package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./src/main/web/views",".html"))

	app.Get("/", func(context iris.Context) {
		context.ViewData("message","hello world")
		context.View("hello.html")
	})

	app.Get("/user/{id:long}", func(context iris.Context) {
		userID,_ := context.Params().GetInt64("id")
		context.Writef("User ID: %d",userID)
	})

	app.Run(iris.Addr(":8080"))
}
