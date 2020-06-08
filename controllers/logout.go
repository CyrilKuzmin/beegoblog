package controllers

import (
	"github.com/astaxie/beego"
)

//LogoutController для очистки сессии
type LogoutController struct {
	beego.Controller
}

//Post для очистки сессии
func (c *LogoutController) Post() {
	sess := c.GetSession("session")
	if sess != nil {
		c.DestroySession()
	} else {
		c.Abort("401")
	}
	c.Redirect("/login", 302)
}
