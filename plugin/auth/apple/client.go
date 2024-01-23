// Package apple SDK
package apple

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"go-zero-yun/plugin"
)

// RespError 微博接口的错误结果返回结构
type RespError struct {
	Error     string `json:"description"`
	ErrorCode int    `json:"error_code"`
}

// HTTPTimeout 请求超时时间 默认 10 秒
var HTTPTimeout time.Duration = time.Second * 10

// apple 实例，在其上实现各类接口
type Apple struct {
	client *http.Client
	cfg    Config
	code   string
}

func New(code string) *Apple {
	// 设置cookiejar后续请求会自动带cookie保持会话
	cfg := GetCfgByName(plugin.DefaultName)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: HTTPTimeout,
	}
	return &Apple{
		client: client,
		cfg:    cfg,
		code:   code,
	}
}
