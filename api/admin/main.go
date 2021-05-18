package main

import (
	"admin/app"
	"admin/middleware/global"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)
var GlobalInfo *global.Global
func main() {
	GlobalInfo = global.NewGlobal()
	defer global.ClearGlobal(GlobalInfo)
	// register struct
	app.RegisterRouter()
	//start all server you register in server/schemas.
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed. " + err.Error())
		return
	}
	chassis.Run()
}