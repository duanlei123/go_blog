package main

import (
	"github.com/beego/i18n"
	"go_blog/models"
	_ "go_blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 设置静态文件
	beego.SetStaticPath("/static","static")
	orm.Debug = true
	_ = orm.RunSyncdb("default", false, true)
	//注册国际化
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	// 注册模板函数
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}

