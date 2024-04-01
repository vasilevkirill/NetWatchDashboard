package routers

import (
	"NetWatchDashboard/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	ui := beego.NewNamespace("/ui",
		beego.NSBefore(controllers.UiAuth),
		beego.NSRouter("/", &controllers.UiController{}),
		beego.NSNamespace("/settings",
			beego.NSBefore(controllers.UiSettingsAuth),
			beego.NSRouter("/users", &controllers.UiSettingsUserController{}),
			beego.NSRouter("/common", &controllers.UiSettingsCommonController{}),
			beego.NSRouter("/devices", &controllers.UiSettingsDevicesController{}),
		),
	)

	beego.AddNamespace(ui)
}
