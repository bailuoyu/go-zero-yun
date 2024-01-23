package reqkit

import (
	"context"
	"github.com/imroc/req/v3"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	"time"
)

const (
	defaultTimeOut = 10 * time.Second
)

// Client 获取客户端
func Client() *req.Client {
	client := req.C()
	if cmdkit.IsEnvDev() {
		client.DevMode()
	}
	//默认请求时间
	client.SetTimeout(defaultTimeOut)
	//设置中间件,记录日志
	client.WrapRoundTripFunc(logkit.ReqLog)
	return client
}

// Req 获取请求
func Req(ctx context.Context) *req.Request {
	client := Client()
	r := client.R().SetContext(ctx)
	return r
}
