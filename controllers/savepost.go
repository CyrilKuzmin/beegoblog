package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/post"
	"github.com/xxlaefxx/beegoblog/models/postdocuments"
	"github.com/xxlaefxx/beegoblog/utils"
)

//SavePostController сохраняет пост
type SavePostController struct {
	beego.Controller
}

//Post сохраняет пост
func (c *SavePostController) Post() {
	id := c.Ctx.Request.FormValue("id")
	title := c.Ctx.Request.FormValue("title")
	content := c.Ctx.Request.FormValue("editor")
	pd := postdocuments.NewPostDocuments()
	if title == "" {
		//Без заголовка не принимаем
		c.Redirect("/error", 302)
		return
	}
	verifyPolicy := utils.MakeNewPolicy()
	if id != "" {
		pd.UpdateOne(post.EditPost(pd.SelectByID(id), title, content, verifyPolicy))
	} else {
		pd.InsertOne(post.NewPost(utils.GenerateUUID(), title, content, verifyPolicy))
	}
	c.Redirect("/blog", 302)
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
