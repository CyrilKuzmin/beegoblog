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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	} else {
		c.Abort("401")
	}
	c.TplName = "lk.html"
	c.Layout = "layout.html"
}
