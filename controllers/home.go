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
	c.TplName = "home.html"
	c.Layout = "layout.html"
}
