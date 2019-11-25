package controllers

import (
	"github.com/astaxie/beego"
	"go_blog/models"
)

// 分类
type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Post() {
	name := this.Input().Get("name")
	err := models.AddCategory(name)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/category", 301)
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	op := this.Input().Get("op")
	switch op {
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0{
			break
		}
		err := models.DelCategory(id)
		if err != nil{
			beego.Error(err)
		}
		this.Redirect("/category", 301)
	}
	this.Data["IsCategory"] = true
	var err error
	this.Data["Categories"], err = models.GetAllCategory()
	this.TplName = "category.tpl"
	if err != nil {
		 beego.Error(err)
	}
}
