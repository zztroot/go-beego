package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int `rom:"pk;auto"`
	Username string
	Password string
}

type Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(30)"`
	Content string `orm:"size(10000)"`
	Times time.Time `orm:"auto_now_add"`
	ArticleType *ArticleType `orm:"rel(fk)"`
}

type ArticleType struct {
	Id int `orm:"pk;auto"`
	TypeName string `orm:"size(30)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init(){
	//&loc=Local 这是控制mysql 数据库的time为当前时间
	orm.RegisterDataBase("default", "mysql", "root:123456789@tcp(localhost:3306)/myblog?charset=utf8&loc=Local")
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}