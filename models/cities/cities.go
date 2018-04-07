package cities

import (
	"DeathStar/models/cities/monaco"
	"log"
)

func InitCities() {
	monaco.InitMonaco()
}

func UpdateCities() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Update failed:", err)
		}
	}()
	monaco.UpdateMonaco()
}
