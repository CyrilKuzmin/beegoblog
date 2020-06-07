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
	c.Data["Posts"] = pdb.SelectAll()
	c.TplName = "blog.html"
	c.Layout = "layout.html"
}
