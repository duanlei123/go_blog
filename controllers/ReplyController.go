package controllers

import (
	"github.com/astaxie/beego"
	"go_blog/models"
)

//评论

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add()  {
	tid := this.Input().Get("tid")
	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")

	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/topic/view/"+tid, 301)
}

func (this *ReplyController) Delete()  {

	if !CheckAccount(this.Ctx){
		return
	}
	tid := this.Input().Get("tid")
	rid := this.Input().Get("rid")
	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/topic/view/"+tid, 301)
}