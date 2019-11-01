package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "quickstart/models"
	_ "quickstart/routers"
)

func main() {
	//beego.AutoRender = false
	//beego.AddFuncMap("showprepage",prepage)
	//beego.AddFuncMap("shownextpage",shownextpage)
	beego.Run()
}
func prepage(pageindex int)(preIndex int){
	preIndex = pageindex - 1
	return
}

func shownextpage(pageindex int)(nextIndex int){
	nextIndex = pageindex + 1
	return
}
