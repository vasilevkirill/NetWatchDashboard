package controllers

import beego "github.com/beego/beego/v2/server/web"

type UiSettingsUserController struct {
	beego.Controller
}

func (c *UiSettingsUserController) Get() {
	c.Data["title"] = "Пользвоатели"
	c.TplName = "settings/users.tpl"
}
