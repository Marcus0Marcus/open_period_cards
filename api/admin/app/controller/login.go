package controller

import (
	"admin/app/protocol"
	"admin/app/service"
	"admin/middleware/constant"
	"admin/middleware/response"
	"admin/middleware/util"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

type LoginCtrl struct {
}

func (r *LoginCtrl) Login(b *restful.Context) {
	req := &protocol.LoginInfoReq{}
	_ = b.ReadEntity(req)
	cond := &service.AdminInfo{
		Phone: req.Phone,
	}
	err, adminInfo := service.NewAdminService().GetAdminByCond(cond)
	if err != nil {
		response.Fail(constant.ErrDb, b)
		return
	}
	if util.Md5(req.Password+adminInfo.Salt) != adminInfo.Pwd {
		response.Fail(constant.ErrPwd, b)
		return
	}
	err, cookie := util.GenLoginCookie(adminInfo.Phone)
	if err != nil {
		response.Fail(err, b)
	}
	util.SetLoginCookie(cookie, b)
	// TODO add redis cookie
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
	}
	// TODO remove redis cookie
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
