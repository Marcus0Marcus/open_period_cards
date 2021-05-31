package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
	"open_period_cards/api/merchant/app/protocol"
	"open_period_cards/data_service"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/response"
	"open_period_cards/middleware/util"
)

type RegisterCtrl struct {
}

func (r *RegisterCtrl) Register(b *restful.Context) {
	req := &protocol.RegisterInfoReq{}
	_ = b.ReadEntity(req)
	cond := &data_service.MerchantInfo{
		Phone: req.Phone,
	}

	err, _ := data_service.NewMerchantService().GetMerchantByCond(cond)
	if err != nil && err.Ret != constant.ErrDBNoRecord.Ret {
		response.Fail(constant.ErrDb, b)
		return
	}
	// 已经注册过了
	if err == nil {
		response.Fail(constant.ErrReReg, b)
		return
	}
	salt := util.GenRandSalt()
	merchantInfo := &data_service.MerchantInfo{
		Phone:  req.Phone,
		Pwd:    util.Md5(req.Password + salt),
		Salt:   salt,
		Status: constant.MerchantStatusApplied,
	}
	err, merchantInfo = data_service.NewMerchantService().CreateMerchant(merchantInfo)
	if err != nil {
		response.Fail(err, b)
		return
	}
	response.Success(b)
}

func (r *RegisterCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/register",
			ResourceFunc: r.Register,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
