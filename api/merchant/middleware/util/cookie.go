package util

import (
	"encoding/base64"
	"fmt"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"merchant/middleware/aes"
	"merchant/middleware/cachehelper"
	"merchant/middleware/constant"
	"merchant/middleware/global"
	"merchant/middleware/response"
	"strings"
	"time"
)

func GenLoginCookie(info string) (*response.FWError, string) {
	uInfo := fmt.Sprintf("%s|%d|%s", global.GetConfig().Config.Cookie.Prefix, time.Now().Unix(), info)
	byteEn, err := aes.AesEncrypt([]byte(uInfo), []byte(global.GetConfig().Config.Cookie.EncryptKey))
	if err != nil {
		return constant.ErrEnc, ""
	}
	return nil, base64.StdEncoding.EncodeToString(byteEn)
}
func ReverseLoginCookie(cookie string) (*response.FWError, string) {
	byteAes, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return constant.ErrEnc, ""
	}
	oriByte, err := aes.AesDecrypt(byteAes, []byte(global.GetConfig().Config.Cookie.EncryptKey))
	if err != nil {
		return constant.ErrEnc, ""
	}
	oriStr := string(oriByte)
	arrStr := strings.Split(oriStr, "|")
	if len(arrStr) != 3 {
		return constant.ErrEnc, ""
	}
	return nil, arrStr[2]
}
func GetCookie(name string, b *restful.Context) (*response.FWError, string) {
	cookie, err := b.Req.Request.Cookie(name)
	if err != nil {
		return constant.ErrLogin, ""
	}
	openlog.Debug(name + cookie.Value)
	return nil, cookie.Value
}
func SetCookie(name string, cookie string, b *restful.Context) {
	b.Resp.AddHeader(name, cookie)
}
func GetLoginCookie(b *restful.Context) (*response.FWError, string) {
	return GetCookie(global.GetConfig().Config.Cookie.Name, b)
}
func SetLoginCookie(cookie string, b *restful.Context) {
	SetCookie(global.GetConfig().Config.Cookie.Name, cookie, b)
}

func GetLoginPhone(b *restful.Context) (*response.FWError, string) {
	// check cookie
	err, cookie := GetLoginCookie(b)
	if err != nil {
		return err, ""
	}
	err, phone := ReverseLoginCookie(cookie)
	if err != nil {
		return err, ""
	}
	// check cache
	err, _ = cachehelper.KeyGet(global.GetConfig().Config.Cache.CookiePrefix + phone)
	if err != nil {
		return constant.ErrLogin, ""
	}
	return nil, phone
}
