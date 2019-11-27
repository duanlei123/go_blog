package routers

import (
	"github.com/astaxie/beego"
	"go_blog/controllers"
	"os"
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
    // 文章
    beego.Router("/topic", &controllers.TopicController{})
    // 文章添加/详情/(自动路由)
    beego.AutoRouter(&controllers.TopicController{})
    // 添加评论
	beego.Router("/reply", &controllers.ReplyController{})
    beego.Router("/reply/add", &controllers.ReplyController{},"post:Add")
    beego.Router("/reply/delete", &controllers.ReplyController{},"get:Delete")

    //创建附件目录
    os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	//beego.SetStaticPath("/attachment","attachment")
    //作为单独的控制器处理
    beego.Router("/attachment/:all",&controllers.AttachmentController{})

}
