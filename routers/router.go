package routers

import (
	"DeathStar/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/parkings", &controllers.ParkingController{})

}
