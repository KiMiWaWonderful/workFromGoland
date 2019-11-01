package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound,notFound)
	app.OnErrorCode(iris.StatusInternalServerError,internalServerError)
	app.Get("/",index)
	app.Run(iris.Addr(":8080"))
}

func notFound(ctx iris.Context)  {
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context)  {
	ctx.WriteString("Oups something went wrong, try again")
}

func index(ctx context.Context)  {
	ctx.View("index.html")
}
