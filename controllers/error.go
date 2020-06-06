package controllers

import (
	"github.com/astaxie/beego"
)

//ErrorController отображает кастомную ошибку
type ErrorController struct {
	beego.Controller
}

//Get возвращает кастомную ошибку
func (c *ErrorController) Get() {
	c.TplName = "error.tpl"
	c.Layout = "layout.tpl"
}
