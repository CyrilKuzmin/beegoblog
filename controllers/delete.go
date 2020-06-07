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
	post := pdb.SelectByID(c.GetString("id"))
	if post != nil {
		pdb.DeleteByID(c.GetString("id"))
	} else {
		c.Redirect("/error", 302)
	}
	c.Redirect("/blog", 302)
}
