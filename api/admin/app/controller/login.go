package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

type LoginCtrl struct {

}
func (r *LoginCtrl) Login(b *restful.Context) {
	phone, err := b.ReadBodyParameter("phone")
	if err != nil {
		b.Write([]byte("get phone: err"))
	}
	b.Write([]byte("get phone: " + phone))
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