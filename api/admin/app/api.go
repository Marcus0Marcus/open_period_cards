package app

import (
	"admin/app/controller"
	"github.com/go-chassis/go-chassis/v2"
)

func RegisterRouter() {
	chassis.RegisterSchema("rest", &controller.RestFulHello{})
	chassis.RegisterSchema("rest", &controller.LoginCtrl{})
}