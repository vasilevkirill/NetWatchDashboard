package controllers

import (
	"github.com/beego/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "My Simple Beego Site"
	c.Data["Email"] = "example@example.com"
	c.TplName = "index.tpl"
}
