package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type UiController struct {
	beego.Controller
}

var UiAuth = func(ctx *context.Context) {
	session, _ := ctx.Session()
	username := session.Get(ctx.Request.Context(), "username")

	if username == nil || username == "" {
		ctx.Redirect(301, "/")
		return
	}

	return
}

func (c *UiController) Get() {
	c.Data["title"] = "UI"
	c.TplName = "index.tpl"
}
