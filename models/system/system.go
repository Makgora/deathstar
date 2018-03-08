package system

import (
	"DeathStar/models/cities"
	"time"
	"log"
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
			//log.Println("Data Update")
			cities.UpdateCities()
			<-time.NewTimer(time.Second * 5).C
		}
	}()
}