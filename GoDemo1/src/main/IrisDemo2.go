package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	users := app.Party("/users",myAuthMiddlewareHandler)
	users.Get("/{id:int}/profile",userProfileHandler)
	users.Get("/inbox/{id:int}",userMessageHandler)
}

func myAuthMiddlewareHandler(ctx iris.Context)  {
		ctx.WriteString("Authentication failed")
}

func userProfileHandler(ctx iris.Context)  {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}

func userMessageHandler(ctx iris.Context)  {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
