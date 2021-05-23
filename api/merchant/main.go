package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
	"merchant/app"
	"merchant/middleware/global"
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
