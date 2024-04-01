package controllers

import beego "github.com/beego/beego/v2/server/web"

type UiSettingsDevicesController struct {
	beego.Controller
}

func (c *UiSettingsDevicesController) Get() {
	c.Data["title"] = "Устройства"
	c.TplName = "settings/devices.tpl"
}
