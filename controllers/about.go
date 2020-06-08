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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	}
	c.TplName = "about.html"
	c.Layout = "layout.html"
}
