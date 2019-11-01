package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/services"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type UserController struct {
	Ctx iris.Context
	Service services.UserService
	Session *sessions.Session
}

const userIDKey  = "UserID"

func (c *UserController) getCurrentUserID() int64  {
	userID := c.Session.GetInt64Default(userIDKey,0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UserController) logout() {
	c.Session.Destroy()
}

var registerStaticView = mvc.View{
	Name:"user/register.html",
	Data:iris.Map{"Title":"User Registration"},
}

func (c *UserController) GetRegister() mvc.Result{
	if c.isLoggedIn(){
		c.logout()
	}
	return registerStaticView
}

func (c *UserController) PostRegister() mvc.Result{
	var (
		firstname = c.Ctx.FormValue("firstname")
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, err := c.Service.Create(password,datamodels.User{
		Username:username,
		Firstname:firstname,
	})

	c.Session.Set(userIDKey,u.ID)
	
	return mvc.Response{
		Err:err,
		Path:"/user/me",
	}
}

var loginStaticView  = mvc.View{
	Name:"user/login.html",
	Data:iris.Map{"Title":"User Login"},
}

func (c *UserController) GetLogin() mvc.Result{
	if c.isLoggedIn(){
		c.logout()
	}
	return loginStaticView
}

func (c *UserController) PostLogin() mvc.Result{
	var(
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, found := c.Service.GetByUsernameAndPassword(username,password)

	if !found {
		return mvc.Response{
			Path:"/user/register",
		}
	}

	c.Session.Set(userIDKey,u.ID)

	return mvc.Response{
		Path:"/user/me",
	}
}

func (c *UserController) GetMe() mvc.Result {
	if !c.isLoggedIn(){
		return mvc.Response{
			Path:"/user/login",
		}
	}

	u,found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		c.logout()
		return c.GetMe()
	}

	return mvc.View{
		Name:"user/me.html",
		Data:iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

func (c *UserController) AnyLogout() {
	if c.isLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/user/login")
}