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

type AccountCtrl struct {
}

func (r *AccountCtrl) Info(b *restful.Context) {
	err, cookie := util.GetLoginCookie(b)
	if err != nil {
		response.Fail(err, b)
		return
	}

	err, phone := util.ReverseLoginCookie(cookie)
	if err != nil {
		response.Fail(err, b)
		return
	}
	cond := &data_service.MerchantInfo{
		Phone: phone,
	}

	err, merchantInfo := data_service.NewMerchantService().GetMerchantByCond(cond)
	if err != nil {
		response.Fail(constant.ErrDb, b)
		return
	}
	resp := &protocol.AccountInfoRes{
		Id:           merchantInfo.Id,
		Avatar:       "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Phone:        merchantInfo.Phone,
		IndustryName: merchantInfo.IndustryName,
		ShopName:     merchantInfo.ShopName,
	}
	response.Data(resp, b)
}

func (r *AccountCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/account/info",
			ResourceFunc: r.Info,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
