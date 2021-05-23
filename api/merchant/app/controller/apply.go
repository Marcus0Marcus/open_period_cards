package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"merchant/app/protocol"
	"merchant/app/service"
	"merchant/middleware/constant"
	"merchant/middleware/response"
	"merchant/middleware/util"
	"net/http"
	"strconv"
)

type ApplyCtrl struct {
}

func (r *ApplyCtrl) Apply(b *restful.Context) {
	// get login phone info
	err, phone := util.GetLoginPhone(b)
	if err != nil {
		response.Fail(constant.ErrLogin, b)
		return
	}
	openlog.Debug(phone)
	// get merchant info
	cond := &service.MerchantInfo{
		Phone: phone,
	}
	err, merchantInfo := service.NewMerchantService().GetMerchantByCond(cond)
	if err != nil {
		response.Fail(constant.ErrLogin, b)
		return
	}
	openlog.Debug(strconv.Itoa(int(merchantInfo.Id)))
	// read request
	req := &protocol.ApplyInfoReq{}
	_ = b.ReadEntity(req)
	openlog.Debug(req.ShopName)
	openlog.Debug(req.IndustryName)
	// fill in merchant struct
	merchantInfo.IndustryName = req.IndustryName
	merchantInfo.ShopName = req.ShopName
	err, _ = service.NewMerchantService().UpdateMerchant(merchantInfo)
	if err != nil {
		response.Fail(err, b)
	}
	response.Success(b)
}

func (r *ApplyCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/apply",
			ResourceFunc: r.Apply,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
