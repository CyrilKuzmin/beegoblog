package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/xxlaefxx/beegoblog/models/user"
)

//LoginController отображает страничку входа
type LoginController struct {
	beego.Controller
}

//Get возвращает страничку входа
func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Layout = "layout.html"
}

//Выполняем проверку email+password
func isLoginDataValid(email, password string) bool {
	u := user.User{Email: email, Password: password}
	err := psqlOrm.Read(&u, "email", "password")
	if err != orm.ErrNoRows {
		return true
	}
	return false
}

//Post возвращает страничку входа
func (c *LoginController) Post() {
	email := c.Ctx.Request.FormValue("email")
	password := c.Ctx.Request.FormValue("password")
	if isLoginDataValid(email, password) {
		//Все ок
		c.Redirect("/blog", 302)
	} else {
		c.Data["Message"] = "Неверный email или пароль"
		c.TplName = "login.html"
		c.Layout = "layout.html"
	}
}
