package controllers

import (
	"github.com/astaxie/beego"
	"DeathStar/models/objects"
)

type ParkingController struct {
	beego.Controller
}

func (c *ParkingController) Get() {
	c.Data["MonacoTotal"] = objects.GetLiveDB().GetCity("Monaco").GetSpacesCount()
	c.Data["MonacoFree"] = objects.GetLiveDB().GetCity("Monaco").GetFreeSpacesCount()
	c.Data["MonacoOcc"] = objects.GetLiveDB().GetCity("Monaco").GetOccSpacesCount()
	lines := objects.GetLiveDB().GetCity("Monaco").String()
	c.Data["Parkings"] = lines
	c.TplName = "parkings.tpl"
}