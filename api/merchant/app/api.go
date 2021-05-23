package app

import (
	"github.com/go-chassis/go-chassis/v2"
	"merchant/app/controller"
)

func RegisterRouter() {
	chassis.RegisterSchema("rest", &controller.LoginCtrl{})
	chassis.RegisterSchema("rest", &controller.RegisterCtrl{})
	chassis.RegisterSchema("rest", &controller.ApplyCtrl{})

}
