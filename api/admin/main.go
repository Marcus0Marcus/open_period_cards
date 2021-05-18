package main

import (
	"admin/app"
	"admin/middleware/global"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)
var globalInfo *global.Global


func main() {
	globalInfo = global.NewGlobal()
	defer global.ClearGlobal(globalInfo)
	// register struct
	app.RegisterRouter()
	//start all server you register in server/schemas.
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed. "+err.Error())
		return
	}
	chassis.Run()
}