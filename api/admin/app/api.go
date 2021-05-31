package app

import (
	"github.com/go-chassis/go-chassis/v2"
	"open_period_cards/api/admin/app/controller"
)

func RegisterRouter() {
	chassis.RegisterSchema("rest", &controller.RestFulHello{})
	chassis.RegisterSchema("rest", &controller.LoginCtrl{})
}
