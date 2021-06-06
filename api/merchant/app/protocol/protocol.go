package protocol

type LoginInfoReq struct {
	Name     string `json:"name"`
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

// 账号信息
type AccountInfoRes struct {
	Id           uint64 `json:"id"`
	Avatar       string `json:"avatar"`
	Phone        string `json:"phone"`
	ShopName     string `json:"shop_name"`
	IndustryName string `json:"industry_name"`
}

// 用户列表
type UserInfoListReq struct {
	PageSize int64 `json:"pageSize"`
	PageNo   int64 `json:"current"`
}
type UserInfoRes struct {
	Id    uint64 `json:"key"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type UserInfoListRes struct {
	Total    int64          `json:"total"`
	PageSize int64          `json:"pageSize"`
	PageNo   int64          `json:"current"`
	Success  bool           `json:"success"`
	List     []*UserInfoRes `json:"data"`
}

// 卡片配置列表
type CardTypeInfoListReq struct {
	PageSize int64 `json:"pageSize"`
	PageNo   int64 `json:"current"`
}
type CardTypeInfoRes struct {
	Id          uint64 `json:"key"`
	Type        uint32 `json:"type"`
	PeriodTimes uint32 `json:"period_times"`
	TotalTimes  uint32 `json:"total_times"`
	Describe    string `json:"describe"`
}

type CardTypeInfoListRes struct {
	Total    int64              `json:"total"`
	PageSize int64              `json:"pageSize"`
	PageNo   int64              `json:"current"`
	Success  bool               `json:"success"`
	List     []*CardTypeInfoRes `json:"data"`
}
