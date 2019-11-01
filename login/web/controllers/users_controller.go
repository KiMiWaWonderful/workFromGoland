package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/services"
)

type UsersController struct {
	Ctx iris.Context
	Service services.UserService
}

func (c *UsersController) Get() (results []datamodels.User) {
	return c.Service.GetAll()
}

func (c *UsersController) GetBy(id int64) (user datamodels.User, found bool){
	u, found := c.Service.GetByID(id)
	if !found {
		// this message will be binded to the
		// main.go -> app.OnAnyErrorCode -> NotFound -> shared/error.html -> .Message text.
		c.Ctx.Values().Set("message", "User couldn't be found!")
	}
	return u, found // it will throw/emit 404 if found == false.
}

func (c *UsersController) PutBy(id int64) (datamodels.User,error) {
	u := datamodels.User{}
	if err := c.Ctx.ReadForm(&u); err != nil{
		return u,err
	}
	return c.Service.Update(id,u)
}

func (c *UsersController) DeleteBy(id int64) interface{}{
	wasDel := c.Service.DeleteByID(id)
	if wasDel{
		return map[string]interface{}{"deleted":id}
	}
	return iris.StatusBadRequest
}