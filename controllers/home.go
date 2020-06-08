package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
)

//HomeController отображает главную страничку
type HomeController struct {
	beego.Controller
}

//unescape используется для отображения поста в красивом HTML вместо чистого кода HTML
func unescape(x string) interface{} {
	return template.HTML(x)
}

//Get возвращает домашнюю страницу
func (c *HomeController) Get() {
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	}
	c.TplName = "home.html"
	c.Layout = "layout.html"
}
