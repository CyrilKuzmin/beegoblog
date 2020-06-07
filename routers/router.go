package routers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/error", &controllers.ErrorController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/post", &controllers.NewPostController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/delete", &controllers.DeleteController{})
	beego.Router("/savepost", &controllers.SavePostController{})
}
