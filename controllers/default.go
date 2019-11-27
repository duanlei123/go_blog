package controllers

import (
	"github.com/astaxie/beego"
	"go_blog/models"
)

type MainController struct {
	beego.Controller
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
