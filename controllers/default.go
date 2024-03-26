package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.vip"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

func (c *MainController) Prepare() {
	// Проверяем, авторизован ли пользователь
	if !c.IsLogin() {
		// Если пользователь не авторизован, делаем редирект на страницу входа
		c.Redirect("/login", 302)
		return
	}
}
