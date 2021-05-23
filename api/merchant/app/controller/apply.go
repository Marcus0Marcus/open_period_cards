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

type ApplyCtrl struct {
}

func (r *ApplyCtrl) Apply(b *restful.Context) {
	// get login phone info
	err, phone := util.GetLoginPhone(b)
	if err != nil {
		response.Fail(constant.ErrLogin, b)
		return
	}

	// read request
	req := &protocol.ApplyInfoReq{}
	_ = b.ReadEntity(req)

	//cond := &service.MerchantInfo{
	//	Phone: phone,
	//}

	merchantInfo := &service.MerchantInfo{
		Phone: phone,
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
