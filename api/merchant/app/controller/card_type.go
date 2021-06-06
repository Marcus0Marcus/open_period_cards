package controller

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"net/http"
	"open_period_cards/api/merchant/app/protocol"
	"open_period_cards/data_service"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/response"
)

type CardTypeCtrl struct {
}

func (r *CardTypeCtrl) List(b *restful.Context) {
	req := &protocol.CardTypeInfoListReq{}
	_ = b.ReadEntity(req)
	cond := &data_service.CardTypeInfo{
		Deleted: 0,
	}
	var card_types []*data_service.CardTypeInfo
	err, total := data_service.NewCardTypeService().GetPagedCardTypeByCond(cond, &card_types, req.PageNo, req.PageSize)
	if err != nil {
		response.Fail(constant.ErrDb, b)
		return
	}
	cardTypesRes := make([]*protocol.CardTypeInfoRes, len(card_types))
	for i := 0; i < len(card_types); i++ {
		cardTypesRes[i] = &protocol.CardTypeInfoRes{
			Id:          card_types[i].Id,
			Type:        card_types[i].Type,
			PeriodTimes: card_types[i].PeriodTimes,
			TotalTimes:  card_types[i].TotalTimes,
			Describe:    card_types[i].Describe,
		}
	}
	resp := &protocol.CardTypeInfoListRes{
		Total:    total,
		PageSize: req.PageSize,
		PageNo:   req.PageNo,
		List:     cardTypesRes,
		Success:  true,
	}
	response.Json(resp, b)
}

func (r *CardTypeCtrl) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:       http.MethodPost,
			Path:         "/card-type/list",
			ResourceFunc: r.List,
			Returns: []*restful.Returns{
				{Code: 200},
			},
		},
	}
}
