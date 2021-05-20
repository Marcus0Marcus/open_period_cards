package constant

import "admin/middleware/response"

var (
	ErrDb = &response.FWError{
		Ret:     -100000,
		Message: "数据库错误",
	}
	ErrEnc = &response.FWError{
		Ret:     -100001,
		Message: "加解密错误",
	}

	ErrPwd = &response.FWError{
		Ret:     -200000,
		Message: "手机号或密码错误",
	}
	ErrLogin = &response.FWError{
		Ret:     -200001,
		Message: "尚未登录",
	}
)
