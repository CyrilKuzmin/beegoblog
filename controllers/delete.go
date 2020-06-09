package controllers

import (
	"github.com/astaxie/beego"
)

//DeleteController удаляет пост
type DeleteController struct {
	beego.Controller
}

//Get запрос удаляет пост
func (c *DeleteController) Get() {
	//Удалить может только админ (пока)
	if c.Data["isAdmin"] == false {
		c.Abort("403")
	}
	post := pdb.SelectByID(c.GetString("id"))
	if post != nil {
		pdb.DeleteByID(c.GetString("id"))
	} else {
		c.Redirect("/error", 302)
	}
	c.Redirect("/blog", 302)
}
