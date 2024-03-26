package main

import (
	_ "NetWatchDashboard/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
