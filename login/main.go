package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/mvc/login/datasource"
	"github.com/kataras/iris/_examples/mvc/login/repositories"
	"github.com/kataras/iris/_examples/mvc/login/services"
	"github.com/kataras/iris/_examples/mvc/login/web/controllers"
	"github.com/kataras/iris/_examples/mvc/login/web/middleware"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	tmpl := iris.HTML("./web/views",".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmpl)

	app.StaticWeb("/public","./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message",ctx.Values().GetStringDefault("message","The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil{
		app.Logger().Fatalf("error while loading the users: %v", err)
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	users := mvc.New(app.Party("/users"))
	users.Router.Use(middleware.BasicAuth)
	users.Register(userService)
	users.Handle(new(controllers.UsersController))

	sessManager := sessions.New(sessions.Config{
		Cookie:"sessioncookiename",
		Expires:24*time.Hour,
	})
	user := mvc.New(app.Party("/user"))
	user.Register(
		userService,
		sessManager.Start,
		)
	user.Handle(new(controllers.UserController))

	app.Run(
		// Starts the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
