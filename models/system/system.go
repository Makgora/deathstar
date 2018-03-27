package system

import (
	"DeathStar/models/cities"
	"time"
)

func Run() {
	initSystem()
	updateSystem()
}

func initSystem() {
	cities.InitCities()
}

func updateSystem() {
	go func() {
		for {
			cities.UpdateCities()
			<-time.NewTimer(time.Second * 60).C
		}
	}()
}
