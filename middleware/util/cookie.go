package util

import (
	"encoding/base64"
	"fmt"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
	"net/url"
	"open_period_cards/middleware/aes"
	"open_period_cards/middleware/cachehelper"
	"open_period_cards/middleware/constant"
	"open_period_cards/middleware/global"
	"open_period_cards/middleware/response"
	"strings"
	"time"
)

func GenLoginCookie(info string) (*response.FWError, string) {
	uInfo := fmt.Sprintf("%s|%d|%s", global.GetConfig().Config.Cookie.Prefix, time.Now().Unix(), info)
	byteEn, err := aes.AesEncrypt([]byte(uInfo), []byte(global.GetConfig().Config.Cookie.EncryptKey))
	if err != nil {
		return constant.ErrEnc, ""
	}
	return nil, url.QueryEscape(base64.URLEncoding.EncodeToString(byteEn))
}
func ReverseLoginCookie(cookie string) (*response.FWError, string) {
	cookie, err := url.QueryUnescape(cookie)
	if err != nil {
		return constant.ErrEnc, ""
	}
	byteAes, err := base64.URLEncoding.DecodeString(cookie)
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
	SetCookie("Set-Cookie", global.GetConfig().Config.Cookie.Name+"="+cookie, b)
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
