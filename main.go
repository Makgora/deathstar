package main

import (
	"DeathStar/models/system"
	_ "DeathStar/routers"
	"github.com/astaxie/beego"
)

func main() {
	system.Run()
	beego.Run()
}
