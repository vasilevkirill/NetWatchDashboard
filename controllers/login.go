package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["Website"] = "My Simple Beego Site"
	this.Data["Email"] = "example@example.com"
	errorMessage := this.GetSession("error")
	if errorMessage != nil {
		this.Data["error"] = errorMessage
		this.DelSession("error")
	}
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	if username != "admin" || password != "admin" {
		this.SetSession("error", "Не правильно имя пользователя или пароль")
		this.Get()
		return
	}
	this.SetSession("username", username)
	this.Redirect("/ui/", 302)
	return

}
