package controllers

import (
	"github.com/astaxie/beego"
)

//NewPostController отображает страничку с созданием поста
type NewPostController struct {
	beego.Controller
}

//Get возвращает страничку с созданием поста (js-редактор)
func (c *NewPostController) Get() {
	c.TplName = "post.tpl"
	c.Layout = "layout.tpl"
}
