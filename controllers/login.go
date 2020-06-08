package controllers

import (
	"fmt"

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
	sess := c.GetSession("session")
	if sess != nil {
		m := sess.((map[string]interface{}))
		c.Data["UserName"] = m["username"]
	}
	c.TplName = "login.html"
	c.Layout = "layout.html"
}

//Выполняем проверку email+password
func isLoginDataValid(email, password string) string {
	u := user.User{Email: email, Password: password, IsActive: true}
	err := psqlOrm.Read(&u, "email", "password")
	if err != orm.ErrNoRows {
		return u.UserName
	}
	return ""
}

//makeSession делает данные сессии для юзера
func makeSession(username string) map[string]interface{} {
	sess := make(map[string]interface{})
	sess["username"] = username
	if username == "admin" {
		sess["isAdmin"] = true
	} else {
		sess["isAdmin"] = false
	}
	u := user.User{UserName: username}
	err := psqlOrm.Read(&u, "user_name")
	if err != orm.ErrNoRows {
		sess["isActive"] = u.IsActive
	}
	return sess
}

//Post возвращает страничку входа
func (c *LoginController) Post() {
	email := c.Ctx.Request.FormValue("email")
	password := c.Ctx.Request.FormValue("password")
	username := isLoginDataValid(email, password)
	if username != "" {
		c.SetSession("session", makeSession(username))
		newData := c.GetSession("session")
		fmt.Println("NEW SESSION DATA: ", newData)
		c.Redirect("/blog", 302)
	} else {
		c.Data["Message"] = "Неверный email или пароль"
		c.TplName = "login.html"
		c.Layout = "layout.html"
	}
}
