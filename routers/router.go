package routers

import (
	"DeathStar/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/team", &controllers.TeamController{})
	beego.Router("/contact", &controllers.ContactController{})
}
