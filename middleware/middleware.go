package middleware

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

//CheckSession проверяем сессию пользователя и устанавливаем нужные параметры
var CheckSession = func(ctx *context.Context) {
	sess := ctx.Input.Session("session")
	fmt.Println("Session ID:", sess)
	if sess != nil {
		m := sess.((map[string]interface{}))
		ctx.Input.SetData("UserName", m["username"])
		ctx.Input.SetData("isAdmin", m["isAdmin"])
		ctx.Input.SetData("isActive", m["isActive"])
	}
}
