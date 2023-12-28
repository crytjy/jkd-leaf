package auth

import (
	"crypto/md5"
	"fmt"
	"github.com/crytjy/jkd-leaf/conf"
	"sort"
	"strconv"
	"strings"
)

type Sign struct {
	key string
}

// NewSign 创建一个具有指定密钥的新 Sign 实例。
func NewSign() *Sign {
	return &Sign{key: conf.SignKey}
}

// CheckSign 验证签名。
func (s *Sign) CheckSign(params map[string]string) bool {
	sign := params["sign"]
	if sign == "" {
		return false
	}

	return s.GetSign(params) == sign
}

// GetSign 生成签名。
func (s *Sign) GetSign(params map[string]string) string {
	delete(params, "sign")

	// 转换 map 为字符串数组，删除空值
	var arr []string
	for k, v := range params {
		if v != "" {
			arr = append(arr, k+"="+v)
		}
	}

	// 按键名升序排序
	sort.Strings(arr)

	// 生成 URL 格式的字符串
	var newQuery string // "0=1703593386&1=77aa087666fa03ff3630668eb622bbfb&key=PI$V1aLYs5bx5eBB!&4KtHokbpTfi46Hr38hiP3mOG9eCqlZnRCyEq8Eoapes37@"
	for i, v := range arr {
		// 查找等号的位置
		equalIndex := strings.Index(v, "=")
		if equalIndex != -1 {
			// 提取等号后面的子字符串
			result := v[equalIndex+1:]
			newQuery += strconv.Itoa(i) + "=" + result + "&"
		}
	}
	str := newQuery + "key=" + s.key

	return strings.ToUpper(s.md5Encode(str))
}

// md5Encode 生成输入字符串的 MD5 散列。
func (s *Sign) md5Encode(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%X", hash.Sum(nil))
}
