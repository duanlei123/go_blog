package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get()  {
	isExit := this.Input().Get("exit") == "true"
	if isExit{
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "no"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd{
		//通过登录验证
		maxAge := 0
		if autoLogin{
			maxAge = 1<<31-1
		}
		//设置cookie
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	//重定向
	this.Redirect("/",301)
	return
}

// 判断用户是否登录
func CheckAccount(ctx *context.Context) bool{
	ck, err := ctx.Request.Cookie("uname")
	if err != nil{
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil{
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}