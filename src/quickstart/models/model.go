package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	Name string `orm:"unique"`
	Pwd string
	Article []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id int
	Aname string
	Atime time.Time `orm:"auto_now"`
	Acount int
	Acontent string
	Aimg string
	Atype string

	ArticleType*ArticleType `orm:"rel(fk)"`
	User []*User `orm:"reverse(many)"`
}

//类型表
type ArticleType struct {
	Id int
	Tname string
	Article []*Article `orm:"reverse(many)"`
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:323900llmT@tcp(localhost:3306)/orm_test?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}
