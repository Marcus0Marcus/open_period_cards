package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/http"
	"open_period_cards/api/merchant/app/protocol"
	"open_period_cards/data_service"
	"open_period_cards/middleware/cachehelper"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/global"
	"open_period_cards/middleware/response"
	"open_period_cards/middleware/util"
)

type LoginCtrl struct {
}

func (r *LoginCtrl) Login(b *restful.Context) {
	req := &protocol.LoginInfoReq{}
	_ = b.ReadEntity(req)
	cond := &data_service.MerchantInfo{
		ShopName: req.Name,
	}

	err, merchantInfo := data_service.NewMerchantService().GetMerchantByCond(cond)
	if err != nil {
		response.Fail(constant.ErrDb, b)
		return
	}
	if util.Md5(req.Password+merchantInfo.Salt) != merchantInfo.Pwd {
		response.Fail(constant.ErrNamePwd, b)
		return
	}
	err, cookie := util.GenLoginCookie(merchantInfo.Phone)
	if err != nil {
		response.Fail(err, b)
	}
	util.SetLoginCookie(cookie, b)
	// add redis record
	err = cachehelper.KeySet(global.GetConfig().Config.Cache.CookiePrefix+merchantInfo.Phone, cookie)
	if err != nil {
		response.Fail(constant.ErrCacheSet, b)
		return
	}
	response.Success(b)
}

func (r *LoginCtrl) Logout(b *restful.Context) {
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
	// check redis exist
	err, _ = cachehelper.KeyGet(global.GetConfig().Config.Cache.CookiePrefix + phone)
	if err != nil {
		openlog.Info(global.GetConfig().Config.Cache.CookiePrefix + phone + " not exist.")
		response.Fail(constant.ErrLogin, b)
		return
	}
	// remove redis
	err = cachehelper.KeyDel(global.GetConfig().Config.Cache.CookiePrefix + phone)
	if err != nil {
		openlog.Info(global.GetConfig().Config.Cache.CookiePrefix + phone + " del err.")
		response.Fail(constant.ErrLogout, b)
	}
	response.Data(phone, b)
}

func (r *LoginCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/login",
			ResourceFunc: r.Login,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
		{
			Method:       http.MethodPost,
			Path:         "/logout",
			ResourceFunc: r.Logout,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
