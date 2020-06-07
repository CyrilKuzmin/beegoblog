package controllers

import (
	"github.com/astaxie/beego"
)

//LoginController отображает страничку входа
type LoginController struct {
	beego.Controller
}

//Get возвращает страничку входа
func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Layout = "layout.html"
}
