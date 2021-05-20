package util

import (
	"admin/middleware/aes"
	"admin/middleware/constant"
	"admin/middleware/response"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

var adminKey = "PeriodCardsAdmin"

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GetAdminCookie(info string) (*response.FWError, string) {
	uInfo := fmt.Sprintf("Admin|%d|%s", time.Now().Unix(), info)
	byteEn, err := aes.AesEncrypt([]byte(uInfo), []byte(adminKey))
	if err != nil {
		return constant.ErrEnc, ""
	}
	return nil, base64.StdEncoding.EncodeToString(byteEn)
}
func ReverseAdminCookie(cookie string) (*response.FWError, string) {
	byteAes, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return constant.ErrEnc, ""
	}
	oriByte, err := aes.AesDecrypt(byteAes, []byte(adminKey))
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
