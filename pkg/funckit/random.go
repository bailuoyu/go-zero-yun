package funckit

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// RandomStr 生成随机字符串
func RandomStr(ln int, ty int) string {
	var chars string
	switch ty {
	case 1:
		chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	case 2:
		chars = "0123456789abcdef"
	case 3:
		chars = "abcdefghijklmnopqrstuvwxyz"
	default:
		chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	bytes := []byte(chars)
	x := int64(len(bytes))
	var result []byte
	for i := 0; i < ln; i++ {
		inx, _ := rand.Int(rand.Reader, big.NewInt(x))
		result = append(result, bytes[inx.Int64()])
	}
	return string(result)
}

func RandomTimeStr(ln int, ty int) string {
	t := time.Now().Unix()
	str := RandomStr(ln, ty)
	return fmt.Sprintf("%d-%s", t, str)
}
