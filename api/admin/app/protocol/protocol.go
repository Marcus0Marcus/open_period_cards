package protocol

type LoginInfoReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
