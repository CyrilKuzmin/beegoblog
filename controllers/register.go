package controllers

import (
	"github.com/astaxie/beego"
)

//RegisterController служит для регистрации новых пользователей
type RegisterController struct {
	beego.Controller
}

//Get возвращает страничку регистрации
func (c *RegisterController) Get() {
	c.TplName = "register.html"
	c.Layout = "layout.html"
}

//Post проверяет данные регистрации и регистрирует пользователя
func (c *RegisterController) Post() {
	c.TplName = "register.html"
	c.Layout = "layout.html"
}
