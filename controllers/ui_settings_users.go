package controllers

import (
	"NetWatchDashboard/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UiSettingsUserController struct {
	beego.Controller
}

func (c *UiSettingsUserController) Get() {
	users, _ := models.GetAllUsers()

	c.Data["users"] = users
	c.Data["title"] = "Пользвоатели"
	c.TplName = "settings/users.tpl"
}
