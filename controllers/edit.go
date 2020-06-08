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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	} else {
		c.Abort("403")
	}
	post := pdb.SelectByID(c.GetString("id"))
	c.Data["Post"] = &post
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
