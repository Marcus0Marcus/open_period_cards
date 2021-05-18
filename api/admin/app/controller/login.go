package controller

import (
	"admin/middleware/response"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

type LoginCtrl struct {

}
type LoginInfoReq struct {
	Phone string    `json:"phone"`
	Password string   `json:"password"`
}
type LoginInfoResp struct {
	Cookie string   `json:"cookie"`
}
func (r *LoginCtrl) Login(b *restful.Context) {
	req := &LoginInfoReq{}
	_ = b.ReadEntity(req)
	resp := &LoginInfoResp{
		Cookie:"ni,a",
	}
	response.Data(resp,b)
}
func (r *LoginCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method: http.MethodPost,
			Path: "/login",
			ResourceFunc: r.Login,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}