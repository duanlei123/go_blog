package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"go_blog/models"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare(){
	lang := this.GetString("lang")
	if lang == "zh-CN"{
		this.Lang = lang
	}else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
	if err != nil{
		beego.Error(err)
	}else {
		c.Data["Topics"] = topics
	}
	//获取所有分类
	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
	c.TplName = "index.tpl"
}
