package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/postdocuments"
)

//BlogController отображает страничку блога
type BlogController struct {
	beego.Controller
}

//Get возвращает страничку блога
func (c *BlogController) Get() {
	c.Data["Posts"] = postdocuments.NewPostDocuments().SelectAll()
	c.TplName = "blog.html"
	c.Layout = "layout.html"
}
