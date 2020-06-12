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
	//Только зарегистрированные могут что-то постить
	if c.Data["UserName"] == nil {
		c.Abort("403")
	}
	c.TplName = "editor.html"
	c.Layout = "layout.html"
}
