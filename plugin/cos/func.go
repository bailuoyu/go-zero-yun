package cos

import (
	"crypto/rand"
	"math/big"
	"net/url"
	"strings"
)

// randomStr 生成随机字符串
func randomStr(ln int) string {
	var chars string
	chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := []byte(chars)
	x := int64(len(b))
	var result []byte
	for i := 0; i < ln; i++ {
		inx, _ := rand.Int(rand.Reader, big.NewInt(x))
		result = append(result, b[inx.Int64()])
	}
	return string(result)
}

// GetKey 获取url中的key值
func GetKey(OriUrl string) (string, error) {
	u, err := url.Parse(OriUrl)
	if err != nil {
		return "", err
	}
	key := strings.TrimLeft(u.Path, "/")
	return key, nil
}
