package main

import (
	_ "NetWatchDashboard/routers"
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

func main() {
	err := prepareRunApplication()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	beego.Run()
}
