package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachmentController struct {
	beego.Controller
}


func (this *AttachmentController) Get()  {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil{
		this.Ctx.WriteString(err.Error())
		return
	}
	file, err := os.Open(filePath)
	if err != nil{
		this.Ctx.WriteString(err.Error())
		return
	}
	defer  file.Close()
	_, err = io.Copy(this.Ctx.ResponseWriter, file)
	if err != nil{
		this.Ctx.WriteString(err.Error())
		return
	}
}
