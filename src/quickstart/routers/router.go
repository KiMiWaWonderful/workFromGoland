package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/register", &controllers.MainController{})
    beego.Router("/index",&controllers.MainController{},"get:ShowIndex")
	//beego.Router("/addArticle",&controllers.MainController{},"get:ShowAdd")
    beego.Router("/addArticle",&controllers.MainController{},"get:ShowAdd;post:HandleAdd")
	beego.Router("/content",&controllers.MainController{},"get:ShowContent")
	beego.Router("/update",&controllers.MainController{},"get:ShowUpdate;post:HandleUpdate")
	beego.Router("/delete",&controllers.MainController{},"get:HandleDelete")
}
