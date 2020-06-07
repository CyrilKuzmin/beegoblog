package main

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/controllers"
	_ "github.com/xxlaefxx/beegoblog/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
