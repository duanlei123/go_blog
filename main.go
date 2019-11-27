package main

import (
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
	beego.Run()
}

