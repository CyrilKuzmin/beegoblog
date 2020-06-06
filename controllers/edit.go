package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/postdocuments"
)

//EditController отображает страничку редактирования поста
type EditController struct {
	beego.Controller
}

//Get возвращает страничку редактирования поста
func (c *EditController) Get() {
	post := postdocuments.NewPostDocuments().SelectByID(c.GetString("id"))
	c.Data["Post"] = &post
	c.TplName = "post.tpl"
	c.Layout = "layout.tpl"
}
