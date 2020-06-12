package controllers

import (
	"github.com/astaxie/beego"
)

//DeletePostController удаляет пост
type DeletePostController struct {
	beego.Controller
}

//Get запрос удаляет пост
func (c *DeletePostController) Get() {
	//Только зарегистрированные могут что-то удалить
	if c.Data["UserName"] == nil {
		c.Abort("403")
	}
	post := pdb.SelectByID(c.GetString("id"))
	if (c.Data["isAdmin"] == false) && (c.Data["UserName"] != post.Author) {
		c.Abort("403")
	}
	if post != nil {
		pdb.DeleteByID(c.GetString("id"))
	} else {
		c.Redirect("/error", 302)
	}
	c.Redirect("/blog", 302)
}
