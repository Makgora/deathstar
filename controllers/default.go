package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "parked.fr"
	c.Data["Email"] = "contact@parked.fr"
	c.TplName = "index.tpl"
}
