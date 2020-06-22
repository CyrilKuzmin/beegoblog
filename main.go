package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/postgres"
	_ "github.com/lib/pq"
	"github.com/xxlaefxx/beegoblog/controllers"
	_ "github.com/xxlaefxx/beegoblog/routers"
)

func main() {
	//Встроенные сессии Beego
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "postgres://root:root@172.16.1.1/blog?sslmode=disable"
	//Обработка ошибок моим контроллером
	beego.ErrorController(&controllers.ErrorController{})
	//Поехали!
	beego.Run()
}
