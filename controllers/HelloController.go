package controllers

import "github.com/astaxie/beego"

// hello world
type HelloController struct {
	beego.Controller
}
func (c *HelloController) Get(){
	c.Ctx.WriteString("hello beego!")
}