package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/postdocuments"
)

//DeleteController удаляет пост
type DeleteController struct {
	beego.Controller
}

//Get запрос удаляет пост
func (c *DeleteController) Get() {
	post := postdocuments.NewPostDocuments().SelectByID(c.GetString("id"))
	if post != nil {
		pd := postdocuments.NewPostDocuments()
		pd.DeleteByID(c.GetString("id"))
	} else {
		c.Redirect("/error", 302)
	}
	c.Redirect("/blog", 302)
}
