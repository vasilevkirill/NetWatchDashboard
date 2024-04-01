package controllers

import (
	"github.com/beego/beego/v2/server/web/context"
)

var UiSettingsAuth = func(ctx *context.Context) {
	session, _ := ctx.Session()
	username := session.Get(ctx.Request.Context(), "username")

	if username == nil || username == "" {
		ctx.Redirect(301, "/")
		return
	}

	return
}
