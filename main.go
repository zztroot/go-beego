package main

import (
	_ "Tianblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("HandleArticleContent", HandleArticleContent)
	beego.Run()
}

func HandleArticleContent(in string)(out string){
	if len(in) >= 100{
		newIn := []rune(in) //[]rune 取字符串的长度，而不是字符串底层占得字节长度
		str := newIn[0:82] //进行切片处理
		return string(str) + "…"
	}else {
		return in
	}
}

