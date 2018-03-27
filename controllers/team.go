package controllers

import (
	"github.com/astaxie/beego"
)

type TeamController struct {
	beego.Controller
}

func (c *TeamController) Get() {
	c.Data[""] = ""
	c.TplName = "team.tpl"
}
