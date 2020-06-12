package routers

import (
	"github.com/astaxie/beego"
	"github.com/xxlaefxx/beegoblog/controllers"
	"github.com/xxlaefxx/beegoblog/middleware"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/error", &controllers.ErrorController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/new", &controllers.NewPostController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/post", &controllers.PostController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/delete", &controllers.DeletePostController{})
	beego.Router("/savepost", &controllers.SavePostController{})
	beego.Router("/lk", &controllers.LKController{})
	beego.Router("/logout", &controllers.LogoutController{})
	//middleware
	beego.InsertFilter("/*", beego.BeforeExec, middleware.CheckSession)
}
