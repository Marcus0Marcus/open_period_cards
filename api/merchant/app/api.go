package app

import (
	"github.com/go-chassis/go-chassis/v2"
	"open_period_cards/api/merchant/app/controller"
)

func RegisterRouter() {
	chassis.RegisterSchema("rest", &controller.LoginCtrl{})
	chassis.RegisterSchema("rest", &controller.RegisterCtrl{})
	chassis.RegisterSchema("rest", &controller.ApplyCtrl{})
	chassis.RegisterSchema("rest", &controller.UserCtrl{})
	chassis.RegisterSchema("rest", &controller.AccountCtrl{})
	chassis.RegisterSchema("rest", &controller.CardTypeCtrl{})

}
