package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type User struct {
	Username string `json:"username"`
    Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	City string `json:"city"`
	Age int `json:"age"`
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./src/main/web/views",".html").Reload(true))

	// 为特定HTTP错误注册自定义处理程序方法
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != ""{
			ctx.Writef("Internal server error : %s",errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	// context.Handler 类型 每一个请求都会先执行此方法 app.Use(context.Handler)
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Info("Begin request for path : %s",ctx.Path())
		ctx.Next()
	})
	
	app.Done(func(ctx context.Context) {

	})

	// Method POST: http://localhost:8080/decode
	app.Post("/decode", func(ctx context.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and comes from %s",user.Firstname,user.Lastname,user.Age,user.City)
	})

	//Method GET: http://localhost:8080/encode
	app.Get("/encode", func(ctx context.Context) {
		doe := User{
			Username:  "Johndoe",
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}
		ctx.JSON(doe)
	})

	//Method GET: http://localhost:8080/profile/anytypeofstring
	app.Get("/profile/{username:string}",profileByUsername)

	//app.Party 定义路由组  第一个参数 设置路由相同的前缀 第二个参数为中间件
	usersRoutes := app.Party("/users", logThisMiddleware)
	//两个{}只是把相同路由组的放在一个区块，没有其他用特殊含义
	{
		// Method GET: http://localhost:8080/users/42
		// 表示id必须是int类型 最小值为 1
		usersRoutes.Get("/{id:int min(1)}", getUserByID)
		// Method POST: http://localhost:8080/users/create
		usersRoutes.Post("/create", createUser)
	}
	//监听 HTTP/1.x & HTTP/2 客户端在  localhost 端口号8080 设置字符集
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))

}

func logThisMiddleware(ctx iris.Context)  {
	ctx.Application().Logger().Info("Path: %s | IP: %s",ctx.Path(),ctx.RemoteAddr())
	ctx.Next()
}

func profileByUsername(ctx iris.Context)  {
	username := ctx.Params().Get("username")
	ctx.ViewData("Username",username)
	ctx.View("profile.html")
}

func getUserByID(ctx iris.Context)  {
	userID := ctx.Params().Get("id")
	user := User{Username:"username"+ userID}
	ctx.XML(user)
}

func createUser(ctx iris.Context) {
	var user User
	//ctx.ReadForm 格式请求数据 与ctx.ReadJSON相似 不过接收的是 Form请求
	//记住 post 字段取名  Username 结构体字段体
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	//{{ . }} 表示 User 的对象  取User的字段 i.e  {{ .Username }} , {{ .Firstname}} etc..

	//向数据模板传值 当然也可以绑定其他值
	ctx.ViewData("", user)
	//渲染模板 ./web/views/create_verification.html
	ctx.View("create_verification.html")
}
