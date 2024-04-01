package controllers

import (
	"NetWatchDashboard/models"
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	errorMessage := c.GetSession("error")
	if errorMessage != nil {
		c.Data["error"] = errorMessage
		c.DelSession("error")
	}
	c.TplName = "login.tpl"

}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	user, err := models.UserAuth(username, password)

	log.Printf("%#v", user)
	if err != nil {
		log.Println(err)
		c.SetSession("error", "Не правильно имя пользователя или пароль")
		c.Get()
		return
	}
	if user.Id == 0 {
		c.SetSession("error", "Не правильно имя пользователя или пароль")
		c.Get()
		return
	}
	log.Printf("%#v", user)
	c.SetSession("username", user.Login)
	c.Redirect("/ui/", 302)
	return
}
