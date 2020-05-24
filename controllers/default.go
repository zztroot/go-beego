package controllers

import (
	"Tianblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

//首页
func (c *MainController) Get() {
	o := orm.NewOrm()
	var article []models.Article
	_, err := o.QueryTable("Article").All(&article)
	if err != nil{
		log.Println("查询文章错误:", err)
		return
	}

	//实现文章显示顺序
	var newArticle []models.Article
	for i:=1; i<=len(article); i++{
		index := article[len(article)-i]
		newArticle = append(newArticle, index)
	}

	//实现文章类型的文章数据
	var articleType []models.ArticleType
	_, err =o.QueryTable("ArticleType").All(&articleType)
	if err != nil{
		log.Println("文章类型查询失败:", err)
		return
	}

	c.Data["articleType"] = articleType
	c.Data["article"] = newArticle
	c.TplName = "index.html"
}

//显示文章函数
func (c *MainController) ShowArticleGet(){
	//获取文章ID
	id, err := c.GetInt("id")
	if err != nil{
		log.Println("获取文章ID失败:", err)
		c.TplName = "showArticle.html"
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil{
		log.Println("查询数据失败:", err)
		c.TplName = "showArticle.html"
		return
	}
	c.Data["article"] = article
	c.TplName = "showArticle.html"
}

//显示登录
func (c *MainController) LoginGet(){
	name := c.Ctx.GetCookie("userName")

	c.Data["name"] = name

	c.TplName = "login.html"
}

//获取登录数据做判断
func (c *MainController) LoginPost(){
	username := c.GetString("Username")
	password := c.GetString("Password")

	if username == "" || password == ""{
		c.TplName = "login.html"
		c.Data["errMsg"] = "用户名和密码为空，请重新输入！"
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Username = username
	err := o.Read(&user, "username")
	if err != nil{
		log.Println("用户名不存在:", err)
		c.TplName = "login.html"
		c.Data["errMsg"] = "用户名不存在"
		return
	}
	if user.Password != password{
		log.Println("密码输入错误:", err)
		c.TplName = "login.html"
		c.Data["errMsg"] = "密码输入错误"
		return
	}

	//登录成功后设置cookie，保存用户名
	c.Ctx.SetCookie("userName", username, time.Second*3600)


	//设置session
	c.SetSession("userName", username)

	c.Redirect("/addArticle", 302)
}

//显示添加文章界面
func (c *MainController) AddArticleGet(){
	//获取session值
	userName := c.GetSession("userName")
	if userName == nil{
		c.Redirect("/login", 302) //如果没有session， 回到登录界面
		return
	}

	//c.DelSession("userName")//成功进入添加文章界面后，直接删除session

	//获取文章类型
	c.TplName = "addArticle.html"
	o := orm.NewOrm()
	var Types []models.ArticleType
	qu := o.QueryTable("article_type")
	_, err := qu.All(&Types)
	if err != nil{
		log.Println("查询文章类型错误,", err)
		return
	}

	c.Data["Types"] = Types
}

//处理文章添加
func (c *MainController) AddArticlePost(){
	title := c.GetString("title_1")
	articleType := c.GetString("ArticleType")
	content := c.GetString("editor01")

	log.Println("文章名称:", title)
	log.Println("文章内容:", content)
	log.Println("文章类型:", articleType)
	articleTypes, _ := strconv.Atoi(articleType)
	o := orm.NewOrm()

	//添加类型到类型表
	Type := models.ArticleType{}
	Type.Id = articleTypes

	//时间初始化
	var times time.Time
	//times := "2006-01-02-15-04-05"

	//插入文章
	article := models.Article{}
	article.Title = title
	article.ArticleType = &Type
	article.Content = content
	article.Times = times
	_, err := o.Insert(&article)
	if err != nil{
		log.Println("文章添加错误：", err)
		return
	}
	log.Println("文章发布成功")
	c.Redirect("/addArticle", 302)

}

//显示类型添加界面
func (c *MainController) AddTypeGet(){
	name := c.GetSession("userName")
	if name == nil{
		c.Redirect("/addArticle", 302)
		return
	}

	c.TplName = "addType.html"
}

//获取类型数据
func (c *MainController) AddTypePost(){
	Type := c.GetString("addType")
	log.Println(Type)
	o := orm.NewOrm()
	articleType := models.ArticleType{}
	articleType.TypeName = Type
	_, err := o.Insert(&articleType)
	if err != nil{
		log.Println("插入文章类型错误:", err)
		return
	}
	c.Redirect("/addArticle", 302)
}

//实现退出登录，并删除掉session
func (c *MainController) OutLoginGet(){
	c.DelSession("userName")
	c.Redirect("/login", 302)
}

