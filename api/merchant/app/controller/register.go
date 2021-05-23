package controller

import (
	"encoding/json"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"merchant/app/protocol"
	"merchant/app/service"
	"merchant/middleware/cachehelper"
	"merchant/middleware/constant"
	"merchant/middleware/response"
	"merchant/middleware/util"
	"net/http"
)

type RegisterCtrl struct {
}

func (r *RegisterCtrl) Register(b *restful.Context) {
	req := &protocol.RegisterInfoReq{}
	_ = b.ReadEntity(req)
	cond := &service.MerchantInfo{
		Phone: req.Phone,
	}

	err, _ := service.NewMerchantService().GetMerchantByCond(cond)
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
	merchantInfo := &service.MerchantInfo{
		Phone:  req.Phone,
		Pwd:    util.Md5(req.Password + salt),
		Salt:   salt,
		Status: constant.MerchantStatusApplied,
	}
	err, merchantInfo = service.NewMerchantService().CreateMerchant(merchantInfo)
	if err != nil {
		response.Fail(err, b)
		return
	}
	// add redis info
	info, _ := json.Marshal(merchantInfo)

	err = cachehelper.KeySet(merchantInfo.Phone, string(info))
	if err != nil {
		openlog.Info(string(info) + " set fail")
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
