package controllers

import (
	"github.com/astaxie/beego"
)

//EditController отображает страничку редактирования поста
type EditController struct {
	beego.Controller
}

//Get возвращает страничку редактирования поста
func (c *EditController) Get() {
	//Только зарегистрированные могут что-то постить
	if c.Data["UserName"] == nil {
		c.Abort("403")
	}
	post := pdb.SelectByID(c.GetString("id"))
	c.Data["Post"] = &post
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
