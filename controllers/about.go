package controllers

import (
	"github.com/astaxie/beego"
)

//AboutController отображает страничку обо мне
type AboutController struct {
	beego.Controller
}

//Get возвращает страничку обо мне
func (c *AboutController) Get() {
	c.TplName = "about.tpl"
	c.Layout = "layout.tpl"
}
