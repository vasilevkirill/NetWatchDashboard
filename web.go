package main

import (
	"github.com/beego/beego"
	"main/controllers"
)

func webRunServer() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.SetStaticPath("/static", "./static")
	beego.BConfig.WebConfig.AutoRender = true
	beego.Run()
}
