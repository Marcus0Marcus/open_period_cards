package protocol

type LoginInfoReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// 注册请求
type RegisterInfoReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// 申请信息
type ApplyInfoReq struct {
	ShopName     string `json:"shop_name"`
	IndustryName string `json:"industry_name"`
}
