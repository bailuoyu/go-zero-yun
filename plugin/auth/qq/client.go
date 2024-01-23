// Package QQ  SDK
package qq

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"go-zero-yun/plugin"
)

// RespError qq接口的错误结果返回结构
type RespError struct {
	Error     string `json:"error_description"`
	ErrorCode int    `json:"error"`
}

type Qq struct {
	client *http.Client
	cfg    Config
	code   string
}

// HTTPTimeout 请求超时时间 默认 10 秒
var HTTPTimeout time.Duration = time.Second * 10

// QQ 实例，在其上实现各类接口
type Weibo struct {
	client       *http.Client
	cfg          Config
	access_token string
}

func New(code string) *Qq {
	cfg := GetCfgByName(plugin.DefaultName)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: HTTPTimeout,
	}
	return &Qq{
		client: client,
		cfg:    cfg,
		code:   code,
	}
}
