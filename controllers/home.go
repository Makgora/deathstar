package controllers

import (
	"DeathStar/models/objects"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type HomeController struct {
	beego.Controller
}

type Parking struct {
	Name       string
	Statut     string
	Address    string
	FreeSpaces string
	Location   Location
}

type Location struct {
	Lat		float32
	Lng		float32
}

//TODO Je dois copier mes modèles (parking par exemple) car je ne peux pas faire {{.parking.GetSpacesCount()}}
func (c *HomeController) Get() {
	c.TplName = "home.tpl"

	city := objects.GetLiveDB().GetCity("Monaco")

	c.Data["cityLocation"] = getCityLocation(city)
	c.Data["parkings"] = getParkings(city)
	c.Data["lastUpdate"] = getLastUpdate(city)
}

func getCityLocation(city *objects.City) *Location {
	return &Location{city.GetLocation().GetLat(), city.GetLocation().GetLng()}
}

func getParkings(city *objects.City) *[]Parking {
	parkings := make([]Parking, 0)
	for _, v := range *city.GetParkings() {
		parkings = append(parkings, Parking{v.GetName(), v.GetStatus(), v.GetAddress(), strconv.Itoa(v.GetFreeSpacesCount()), Location{v.GetLocation().GetLat(), v.GetLocation().GetLng()}})
	}
	return &parkings
}

func getLastUpdate(city *objects.City) int {
	return int(time.Since(city.GetLastUpdate()).Seconds())
}