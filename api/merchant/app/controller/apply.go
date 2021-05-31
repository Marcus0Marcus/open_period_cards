package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
	"open_period_cards/api/merchant/app/protocol"
	"open_period_cards/data_service"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/response"
	"open_period_cards/middleware/util"
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
	cond := &data_service.MerchantInfo{
		Phone: phone,
	}
	err, merchantInfo := data_service.NewMerchantService().GetMerchantByCond(cond)
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
	err, _ = data_service.NewMerchantService().UpdateMerchant(merchantInfo)
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
