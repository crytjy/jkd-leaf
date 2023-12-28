package auth

import (
	"github.com/crytjy/jkd-leaf/conf"
	"github.com/crytjy/jkd-leaf/log"
	"net/url"
)

type ConnParam struct {
	ReqTime int64
	Sign    string
	Token   string
}

func CheckUser(query string) bool {
	if conf.SignKey == "" {
		return true
	}
	params, err := url.ParseQuery(query)
	if err != nil {
		log.Debug("Error parsing query:", err)
		return false
	}

	token := params.Get("token")
	sign := params.Get("sign")
	reqTime := params.Get("reqTime")

	if token == "" || sign == "" || reqTime == "" {
		log.Debug("Error parsing query:", err)
		return false
	}

	return NewSign().CheckSign(map[string]string{
		"token":   token,
		"sign":    sign,
		"reqTime": reqTime,
	})
}
