package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
}

func CheckLogin(c *context.Context) {

}
func (c *MainController) Get() {
	if c.IsLogin() {
		c.Redirect("/ui/", 302)
		return
	}

	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Prepare() {
	// Проверяем, авторизован ли пользователь
	if !c.IsLogin() {
		// Если пользователь не авторизован, делаем редирект на страницу входа
		c.Redirect("/login", 302)
		return
	}
}

func (c *MainController) IsLogin() bool {

	// Получаем значение сессии с ключом "username"
	username := c.GetSession("username")
	// Проверяем, существует ли значение сессии
	if username == nil {
		return false
	}
	// В данном примере просто проверяем, что имя пользователя не пустое
	// В вашем реальном приложении здесь должна быть более сложная логика
	return username != ""

}
