package controllers

import beego "github.com/beego/beego/v2/server/web"

type UiSettingsCommonController struct {
	beego.Controller
}

func (c *UiSettingsCommonController) Get() {
	c.Data["title"] = "Общие настройки"
	c.TplName = "settings/common.tpl"
}
