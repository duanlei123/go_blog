package controllers

import (
	"github.com/astaxie/beego"
	"go_blog/models"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}
//文章页
func (this *TopicController) Get(){
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["IsTopic"] = true
	topics, err := models.GetAllTopics("","", false)
	if err != nil{
		beego.Error(err)
	}else {
		this.Data["Topics"] = topics
	}
	this.TplName = "topic.tpl"
}
func (this *TopicController) Add(){
	this.TplName = "topic_add.tpl"
}
// 添加文章 / 修改文章
func (this *TopicController) Post(){
	// 添加文章，必须管理员添加
	if !CheckAccount(this.Ctx){
		this.Redirect("/login", 302)
		return
	}
	// 获取表单数据
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	labels := this.Input().Get("labels")

	// 获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil{
		beego.Error(err)
	}
	var attachment string
	if fh != nil{
		//保持附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil{
			beego.Error(err)
		}
	}
	if len(tid) == 0{
		err = models.AddTopic(title,category,labels, content,attachment)
	}else {
		err = models.ModelsTopic(tid,title,category,labels,content,attachment)
	}
	if err != nil{
		beego.Error(err)
	}
	this.Redirect("/topic", 301)
}

//文章详情
func (this *TopicController) View(){
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 301)
		return
	}
	//获取评论
	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil{
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels," ")
	this.Data["Tid"] = this.Ctx.Input.Param("0")
	this.TplName = "topic_view.tpl"
}
// 修改文章
func (this *TopicController) Modify(){
	// 添加文章，必须管理员添加
	if !CheckAccount(this.Ctx){
		this.Redirect("/login", 302)
		return
	}
	// 获取文章id
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil{
		beego.Error(err)
		this.Redirect("/", 301)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.TplName = "topic_modify.tpl"
}

func (this *TopicController) Delete(){
	//必须管理员删除
	if !CheckAccount(this.Ctx){
		this.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(this.Ctx.Input.Param("0"))
	if err != nil{
		beego.Error(err)
	}
	this.Redirect("/", 301)
}

