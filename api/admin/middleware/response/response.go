package response

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
)

type Resp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
type RespData struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type FWError struct {
	Ret     int32
	Message string
}

func Json(value interface{}, b *restful.Context) {
	b.WriteJSON(value, "application/json")
}
func Success(b *restful.Context) {
	b.WriteJSON(&Resp{Code: 0, Message: "success"}, "application/json")
}
func Data(data interface{}, b *restful.Context) {
	b.WriteJSON(&RespData{Code: 0, Message: "success", Data: data}, "application/json")
}
func Fail(error *FWError, b *restful.Context) {
	b.WriteJSON(&Resp{Code: error.Ret, Message: error.Message}, "application/json")
}
