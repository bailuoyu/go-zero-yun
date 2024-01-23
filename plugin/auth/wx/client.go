// Package wx wx SDK
package wx

import (
	"net/http"
	"time"

	"go-zero-yun/plugin"
)

// RespError 微博接口的错误结果返回结构
type RespError struct {
	Errmsg    string `json:"errmsg"`
	ErrorCode int    `json:"errcode"`
	Request   string `json:"request"`
}

// HTTPTimeout 请求超时时间 默认 10 秒
var HTTPTimeout time.Duration = time.Second * 10

// WX 实例，在其上实现各类接口
type Wx struct {
	client *http.Client
	cfg    Config
	code   string
}

func New(code string) *Wx {
	cfg := GetCfgByName(plugin.DefaultName)

	client := &http.Client{
		Timeout: HTTPTimeout,
	}
	return &Wx{
		client: client,
		cfg:    cfg,
		code:   code,
	}
}
