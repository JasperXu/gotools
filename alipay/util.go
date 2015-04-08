package alipay

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

var (
	AlipayPartner  string //合作者ID
	AlipayKey      string //合作者私钥
	WebReturnUrl   string //网站同步返回地址
	WebNotifyUrl   string //网站异步返回地址
	WebSellerEmail string //网站卖家邮箱地址
)

var GATEWAY string = "https://mapi.alipay.com/gateway.do" // Alipay 通信网关

func computeSignature(params map[string]string) string {
	var keyStrings []string
	for k, _ := range params {
		if k != "sign" && k != "sign_type" {
			keyStrings = append(keyStrings, k)
		}
	}
	sort.Strings(keyStrings)
	paramsString := ""
	for i := 0; i < len(keyStrings); i++ {
		if i == 0 {
			paramsString = keyStrings[i] + "=" + params[keyStrings[i]]
		}
		paramsString += "&" + keyStrings[i] + "=" + params[keyStrings[i]]
	}

	paramsString += AlipayKey

	//md5加密
	h := md5.New()
	h.Write([]byte(paramsString))
	return hex.EncodeToString(h.Sum(nil))
}

func Verify(params map[string]string, sign string) bool {
	return computeSignature(params) == sign
}
