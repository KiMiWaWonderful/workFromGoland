package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.Use(recover2.New())
	//app.Use(logger.New())

	//for ; ;  {
	//	app.Handle("GET","/", func(context iris.Context) {
	//		context.HTML("<h1>Welcome</h1>")
	//
	//	})

	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog:false}))
	app.Get("/ping", func(context iris.Context) {
		context.WriteString("pong")
	})

	app.Get("/hello", func(context iris.Context) {
		context.JSON(iris.Map{"message":"hello"})
	})

	app.Get("/hello/{name}", func(ctx context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		name := ctx.Params().Get("name")
		ctx.HTML("<h1>"+name+"</h1>")
	})

	app.Get("/api/users/{userid:uint64}", func(context context.Context) {
		userID, err := context.Params().GetUint("userid")

		if err != nil{
			context.JSON(map[string]interface{}{
				"requestcode":  201,
				"message":  "bad request",
			})
			return
		}

		context.JSON(map[string]interface{}{
			"requestcode":  200,
			"message":  userID,
		})
	})

	app.Get("/api/users/{isLogin:bool}", func(i context.Context) {
		isLogin,err := i.Params().GetBool("isLogin")

		if err != nil{
			i.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			i.WriteString("已登录")
		}else {
			i.WriteString("未登录")
		}
	})
	app.Handle("GET","/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.HTML(path)
	})
	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed),iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:false,
		DisablePathCorrection:false,
		EnablePathEscape:false,
		FireMethodNotAllowed:false,
		DisableBodyConsumptionOnUnmarshal:false,
		DisableAutoFireStatusCode:false,
		TimeFormat:"Mon,02 Jan 2006 15:04:05 GMT",
		Charset:"UTF-8",
	}))
	//}

}
