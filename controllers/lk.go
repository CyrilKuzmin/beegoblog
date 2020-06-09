package controllers

import (
	"github.com/astaxie/beego"
)

//LKController отображает страничку личного кабинета
type LKController struct {
	beego.Controller
}

//Get возвращает страничку личного кабинета
func (c *LKController) Get() {
	if c.Data["UserName"] == nil {
		c.Abort("401")
	}
	c.TplName = "lk.html"
	c.Layout = "layout.html"
}
