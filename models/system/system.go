package system

import (
	"DeathStar/models/cities"
)

func Run() {
	initSystem()
	updateSystem()
/*
	for true {
		updateSystem()
		time.Sleep(time.Second * 10)
	}*/
}

func initSystem() {
	cities.InitCities()
}

func updateSystem() {
	cities.UpdateCities()
}