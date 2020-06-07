package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

//BlogController отображает страничку блога
type BlogController struct {
	beego.Controller
}

//Get возвращает страничку блога
func (c *BlogController) Get() {
	fmt.Println("GET ЗАПРОС БЛОГА")
	c.Data["Posts"] = pdb.SelectAll()
	c.TplName = "blog.html"
	c.Layout = "layout.html"
}
