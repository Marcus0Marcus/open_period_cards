package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
	"open_period_cards/api/merchant/app/protocol"
	"open_period_cards/data_service"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/response"
)

type UserCtrl struct {
}

func (r *UserCtrl) List(b *restful.Context) {
	req := &protocol.UserInfoListReq{}
	_ = b.ReadEntity(req)
	cond := &data_service.UserInfo{
		Deleted: 0,
	}
	var users []*data_service.UserInfo
	err, total := data_service.NewUserService().GetPagedUserByCond(cond, &users, req.PageNo, req.PageSize)
	if err != nil {
		response.Fail(constant.ErrDb, b)
		return
	}
	usersRes := make([]*protocol.UserInfoRes, len(users))
	for i := 0; i < len(users); i++ {
		usersRes[i] = &protocol.UserInfoRes{
			Id:    users[i].Id,
			Phone: users[i].Phone,
			Name:  users[i].Name,
		}
	}
	resp := &protocol.UserInfoListRes{
		Total:    total,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
		List:     usersRes,
		Success:  true,
	}
	response.Json(resp, b)
}

func (r *UserCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/user/list",
			ResourceFunc: r.List,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
