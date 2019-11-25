package routers

import (
	"github.com/astaxie/beego"
	"go_blog/controllers"
)

func init() {
	// 注册首页路由
    beego.Router("/", &controllers.MainController{})
    // 测试框架搭建
    beego.Router("/hello", &controllers.HelloController{})
    // 登录
    beego.Router("/login", &controllers.LoginController{})
    // 分类
    beego.Router("/category", &controllers.CategoryController{})

}
