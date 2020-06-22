package controllers

import (
	"github.com/astaxie/beego"
)

//PostController отображает страничку поста
type PostController struct {
	beego.Controller
}

//Get возвращает страничку поста
func (c *PostController) Get() {
	post := pdb.SelectByID(c.GetString("id"))
	c.Data["Post"] = &post
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
