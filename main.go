package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/xxlaefxx/beegoblog/controllers"
	_ "github.com/xxlaefxx/beegoblog/routers"
)

func main() {
	orm.Debug = true
	//Обработка ошибок моим контроллером
	beego.ErrorController(&controllers.ErrorController{})
	//Поехали!
	beego.Run()
}
