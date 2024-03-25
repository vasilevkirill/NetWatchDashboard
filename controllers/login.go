package controllers

import "github.com/beego/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "My Simple Beego Site"
	c.Data["Email"] = "example@example.com"
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {

}
