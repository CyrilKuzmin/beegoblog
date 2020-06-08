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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	} else {
		c.Abort("403")
	}
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
