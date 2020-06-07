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
	c.TplName = "error.html"
	c.Layout = "layout.html"
}

//Error401 обрабатывает ошибки аутентификации
func (c *ErrorController) Error401() {
	c.Data["Content"] = "Вы не авторизованы"
	c.TplName = "error.html"
	c.Layout = "layout.html"
}

//Error403 обрабатывает ошибки авторизации
func (c *ErrorController) Error403() {
	c.Data["Content"] = "Доступ запрещен"
	c.TplName = "error.html"
	c.Layout = "layout.html"
}

//Error404 Просто 404
func (c *ErrorController) Error404() {
	c.Data["Content"] = "Запрашиваемая страница не найдена"
	c.TplName = "error.html"
	c.Layout = "layout.html"
}
