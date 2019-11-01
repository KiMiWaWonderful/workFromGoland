package main

import "github.com/kataras/iris"

func main(){
	app := iris.New()

	h := func(ctx iris.Context) {
		ctx.HTML("<b>Hi</b>")
	}

	home := app.Get("/",h)
	home.Name = "home"

	app.Get("/about",h).Name = "about"
	app.Get("/page/{id}",h).Name = "page"

	app.Run(iris.Addr(":8080"))
}
