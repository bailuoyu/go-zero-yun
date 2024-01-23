// Package weibo 新浪微博 SDK
package weibo

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

// RespError 微博接口的错误结果返回结构
type RespError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
	Request   string `json:"request"`
}

// HTTPTimeout 请求超时时间 默认 10 秒
var HTTPTimeout time.Duration = time.Second * 10

// Weibo 实例，在其上实现各类接口
type Weibo struct {
	client       *http.Client
	access_token string
}

func New(access_token string) *Weibo {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: HTTPTimeout,
	}
	return &Weibo{
		client:       client,
		access_token: access_token,
	}
}
