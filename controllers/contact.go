package controllers

import (
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.Data[""] = ""
	c.TplName = "contact.tpl"
}
