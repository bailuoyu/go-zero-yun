package logkit

import (
	"github.com/imroc/req/v3"
	"go-zero-yun/pkg/funckit"
	"net/http"
)

type ReqParams struct {
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Body   string      `json:"body"`
	Header http.Header `json:"header"`
}

func ReqLog(rt req.RoundTripper) req.RoundTripFunc {
	return func(req *req.Request) (*req.Response, error) {
		//执行请求之前
		resp, err := rt.RoundTrip(req)
		//执行请求之后
		//请求时间
		runtime := funckit.DurToMic2(resp.TotalTime())
		//记录请求和返回日志
		reqp := ReqParams{
			Method: req.Method,
			Url:    req.URL.String(),
			Body:   string(req.Body),
			Header: req.Headers,
		}
		//记录日志
		lgr := WithCallerSkip(2)
		if err != nil {
			lgr.WithType(LogHttp).WithRuntime(runtime).Errorf(req.Context(), "error:%s, req:%+v, resp:%+v", err.Error(), reqp, resp)
		} else {
			lgr.WithType(LogHttp).WithRuntime(runtime).Infof(req.Context(), "success, req:+%v, resp:%+v", reqp, resp)
		}
		return resp, err
	}
}
