package main

import (
	"github.com/beego/beego"
	"main/controllers"
)

func webRunServer() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
