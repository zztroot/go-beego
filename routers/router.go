package routers

import (
	"Tianblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/showArticle", &controllers.MainController{}, "get:ShowArticleGet")
	beego.Router("/login", &controllers.MainController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/addArticle", &controllers.MainController{}, "get:AddArticleGet;post:AddArticlePost")
    beego.Router("/addType", &controllers.MainController{}, "get:AddTypeGet;post:AddTypePost")

    //退出登录路由
    beego.Router("/outLogin", &controllers.MainController{}, "get:OutLoginGet")

}
