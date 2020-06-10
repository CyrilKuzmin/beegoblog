package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/models/post"
	"github.com/xxlaefxx/beegoblog/utils"
)

//SavePostController сохраняет пост
type SavePostController struct {
	beego.Controller
}

func ImageProcessing(content string) string {
	return fmt.Sprintln("Исходные данные", content)
}

//Post сохраняет пост
func (c *SavePostController) Post() {
	//Только зарегистрированные могут что-то постить
	if c.Data["UserName"] == nil {
		c.Abort("403")
	}
	id := c.Ctx.Request.FormValue("id")
	title := c.Ctx.Request.FormValue("title")
	content := c.Ctx.Request.FormValue("editor")
	if title == "" {
		//Без заголовка не принимаем
		c.Redirect("/error", 302)
		return
	}
	fmt.Println(ImageProcessing(content))
	verifyPolicy := utils.MakeNewPolicy()
	if id != "" {
		pdb.UpdateOne(post.EditPost(pdb.SelectByID(id), title, content, verifyPolicy))
	} else {
		pdb.InsertOne(post.NewPost(utils.GenerateID(8), title, content, verifyPolicy))
	}
	c.Redirect("/blog", 302)
	c.TplName = "post.html"
	c.Layout = "layout.html"
}
