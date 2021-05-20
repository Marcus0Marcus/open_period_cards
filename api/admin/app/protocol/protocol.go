package protocol

type LoginInfoReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type LoginInfoResp struct {
	Cookie string `json:"cookie"`
}

type LogoutInfoReq struct {
	Cookie string `json:"cookie"`
}
