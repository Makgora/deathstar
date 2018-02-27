package main

import (
	_ "DeathStar/routers"
	"DeathStar/models/monaco"
)

func main() {
	//beego.Run()
	monaco.UpdateMonaco()
}
