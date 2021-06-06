package constant

import "open_period_cards/middleware/response"

var (
	ErrDb = &response.FWError{
		Ret:     -100000,
		Message: "数据库错误",
	}
	ErrEnc = &response.FWError{
		Ret:     -100001,
		Message: "加解密错误",
	}
	ErrSys = &response.FWError{
		Ret:     -100002,
		Message: "系统错误，请稍后再试",
	}
	ErrDBNoRecord = &response.FWError{
		Ret:     -100003,
		Message: "记录未找到",
	}
	ErrCacheSet = &response.FWError{
		Ret:     -100004,
		Message: "缓存设置失败",
	}
	ErrCacheGet = &response.FWError{
		Ret:     -100005,
		Message: "缓存查询失败",
	}
	ErrCacheDel = &response.FWError{
		Ret:     -100006,
		Message: "缓存删除失败",
	}
	ErrCacheNotExist = &response.FWError{
		Ret:     -100007,
		Message: "缓存不存在",
	}
	ErrMarshal = &response.FWError{
		Ret:     -100008,
		Message: "序列化失败",
	}
	ErrUnMarshal = &response.FWError{
		Ret:     -100009,
		Message: "反序列化失败",
	}

	ErrNamePwd = &response.FWError{
		Ret:     -200000,
		Message: "账号或密码错误",
	}
	ErrLogin = &response.FWError{
		Ret:     -200001,
		Message: "尚未登录",
	}
	ErrReReg = &response.FWError{
		Ret:     -200002,
		Message: "此手机号已注册",
	}
	ErrLogout = &response.FWError{
		Ret:     -200003,
		Message: "退出失败",
	}
)

var (
	MerchantStatusApplied int32 = 0 // 申请已提交
	MerchantStatusPassed  int32 = 1 // 申请成功
	MerchantStatusDenied  int32 = 2 // 申请失败
)
