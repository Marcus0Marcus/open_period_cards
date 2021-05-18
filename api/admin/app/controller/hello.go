package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
)

type RestFulHello struct {

}
func (r *RestFulHello) SayHello(b *restful.Context) {
	b.Write([]byte("get user id: " + b.ReadPathParameter("userid")))
}
func (r *RestFulHello) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method: http.MethodGet,
			Path: "/sayhello/{userid}",
			ResourceFunc: r.SayHello,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}