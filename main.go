package main

import (
	_ "DeathStar/routers"
	"DeathStar/models/system"
	"github.com/astaxie/beego"
)

func main() {
	system.Run()
	beego.Run()
}
