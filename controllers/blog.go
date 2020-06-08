package controllers

import (
	"github.com/astaxie/beego"
)

//BlogController отображает страничку блога
type BlogController struct {
	beego.Controller
}

//Get возвращает страничку блога
func (c *BlogController) Get() {
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	}
	c.Data["Posts"] = pdb.SelectAll()
	c.TplName = "blog.html"
	c.Layout = "layout.html"
}
