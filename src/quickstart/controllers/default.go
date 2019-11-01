package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"quickstart/models"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Name = "111"
	//user.Pwd = "222"
	//o.Insert(&user)

	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 1
	//err := o.Read(&user)
	//if err != nil{
	//	beego.Info("fail",err)
	//	return
	//}
	//beego.Info("user:",user)

	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 1
	//err := o.Read(&user)
	//if err == nil {
	//	user.Name = "222"
	//	user.Pwd = "333"
	//
	//	_,err = o.Update(&user)
	//	if err != nil{
	//		beego.Info("更新失败",err)
	//		return
	//	}
	//}
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	c.TplName = "register.html"
}

func (c *MainController)Post()  {
	userName :=c.GetString("userName")
	pwd := c.GetString("pwd")
	//beego.Info(userName,pwd)
	if userName == "" || pwd == "" {
		beego.Info("不能为空")
		c.Redirect("/register",302)
		return
	}

	o := orm.NewOrm()

	user := models.User{}
	user.Name = userName
	user.Pwd= pwd
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
		c.Redirect("/register",302)
		return
	}
	c.Ctx.WriteString("注册成功")
}

func (c *MainController)ShowLogin()  {
	c.TplName = "login.html"
}

func (c *MainController)HandleLogin()  {
	userName :=c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Info("不能为空")
		//c.Redirect("/login",302)
		c.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil{
		c.TplName = "login.html"
		return
	}
	//c.Ctx.WriteString("welcome")
	c.Redirect("/index",302)
}

func (c *MainController)ShowIndex()  {
	o := orm.NewOrm()
	//id,_ := c.GetInt("select")

	var articles []models.Article
	_,err := o.QueryTable("Article").All(&articles)
	if err != nil{
		beego.Info("查询所有文章信息出错")
		return
	}

	//var count int64
	//
	//if id == 0 {
	//	count,err = o.QueryTable("Article").RelatedSel("ArticleType").Count()
	//}else {
	//	count,err = o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__Id",id).Count()
	//}
	//if err != nil {
	//	beego.Info("查询失败",err)
	//	return
	//}
	//pagesize := int64(2)
	//
	//index, err := c.GetInt("pageIndex")
	//if err != nil{
	//	index = 1
	//}
	//pageCount := math.Ceil(float64(count) / float64(pagesize))   //总页数
	//
	//if index <=0 || index > int(pageCount){
	//	index = 1
	//}
	//
	//start := (int64(index) -1) * pagesize
	//var artis []models.Article
	//if id == 0{
	//	o.QueryTable("Article").Limit(pagesize,start).RelatedSel("ArticleType").All(&artis)
	//}else {
	//	o.QueryTable("Article").Limit(pagesize,start).RelatedSel("ArticleType").Filter("ArticleType__Id",id).All(&artis)
	//}
	//
	////获取类型数据
	//var artiTypes []models.ArticleType
	//_,err = o.QueryTable("ArticleType").All(&artiTypes)
	//if err != nil{
	//	beego.Info("获取类型错误")
	//	return
	//}

	//beego.Info(articles)
	c.Data["articles"] = articles
	//c.Data["pageCount"] = pageCount
	//c.Data["count"] = count
	//c.Data["articles"] = artis
	//c.Data["pageIndex"] = index
	//c.Data["typeid"] = id    //文章类型ID
	c.TplName = "index.html"
}

func (c *MainController)ShowAdd()  {
	c.TplName = "add.html"
}

func (c *MainController)HandleAdd()  {
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	f,h,err := c.GetFile("uploadname")
	defer f.Close()

	fileext := path.Ext(h.Filename)
	if fileext != ".jpg" && fileext != ".png"{
		beego.Info("上传文件格式错误")
		return
	}
	if h.Size > 500000000 {
		beego.Info("上传文件过大")
		return
	}
	filename := time.Now().Format("2006-01-02 15:04:05")+fileext
	if err != nil{
		beego.Info("上传文件失败")
		return
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	if artiContent == "" || artiName == ""{
		beego.Info("添加文章数据错误")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{}
	arti.Aname = artiName
	arti.Acontent = artiContent
	arti.Aimg = "/static/img/"+filename
	_,err = o.Insert(&arti)
	if err != nil{
		beego.Info("插入数据出错")
		return
	}
	c.Redirect("/index",302)
}

func (c *MainController)ShowContent()  {
	id,err := c.GetInt("id")
	if err != nil{
		beego.Info("获取文章ID失败")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误ShowContent",err)
		return
	}
	c.Data["article"] = arti
	c.TplName = "content.html"
}

func (c *MainController)ShowUpdate()  {
	id,err := c.GetInt("id")
	if err != nil{
		beego.Info("获取文章ID失败")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误ShowUpdate",err)
		return
	}
	c.Data["article"] = arti
	c.TplName = "update.html"
}

func (c *MainController)HandleUpdate()  {
	//1.拿到数据
	id,_ := c.GetInt("id")
	artiName := c.GetString("articleName")
	content := c.GetString("content")
	f,h,err:=c.GetFile("uploadname")
	var filename string
	if err != nil{
		beego.Info("上传文件失败")
		return
	}else {
		defer f.Close()
		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != "png"{
			beego.Info("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("2006-01-02 15:04:05") + fileext  //6-1-2 3:4:5
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	//2.对数据进行一个处理
	if artiName == "" || content ==""{
		beego.Info("更新数据获取失败")
		return
	}

	//3.更新操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询数据错误")
		return
	}
	arti.Aname = artiName
	arti.Acontent = content
	arti.Aimg = "./static/img/"+filename

	_,err = o.Update(&arti,"Aname","Acontent","Aimg")
	if err != nil{
			beego.Info("更新数据显示错误")
		return
	}

	//4.返回列表页面
	c.Redirect("/index",302)
}

func (c*MainController)HandleDelete(){
	//1.拿到数据
	id,err:=c.GetInt("id")
	if err != nil{
		beego.Info("获取id数据错误")
		return
	}


	//2.执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误HandleDelete")
		return
	}
	o.Delete(&arti)

	//3.返回列表页面
	c.Redirect("/index",302)
}