package logkit

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/utils"
	"go-zero-yun/pkg/funckit"
	"io"
	"net/http"
	"time"
)

type EsDoer struct {
	Client *http.Client
}

// GetEsDoer 获取elastic的doer
func GetEsDoer() *EsDoer {
	client := http.DefaultClient
	//默认请求时间
	client.Timeout = 20 * time.Second
	return &EsDoer{Client: client}
}

func (dr *EsDoer) Do(req *http.Request) (*http.Response, error) {
	//跳过head日志
	if req.Method == http.MethodHead {
		return dr.Client.Do(req)
	}
	//请求参数
	reqp := ReqParams{
		Url:    req.URL.String(),
		Method: req.Method,
		Header: req.Header,
	}
	if req.Body != nil {
		jb, err := io.ReadAll(req.Body)
		if err != nil {
			reqp.Body = ""
		} else {
			reqp.Body = string(jb)
			req.Body = io.NopCloser(bytes.NewBuffer(jb))
		}
	}
	//计时
	timer := utils.NewElapsedTimer()
	resp, err := dr.Client.Do(req)
	//总时间
	runtime := funckit.DurToMic2(timer.Duration())
	//记录日志
	lgr := WithCallerSkip(2)
	if err != nil {
		lgr.WithType(LogES).WithRuntime(runtime).
			Errorf(req.Context(), "error:%s, req:%+v, code:%d", err.Error(), reqp, resp.StatusCode)
	} else {
		lgr.WithType(LogES).WithRuntime(runtime).
			Infof(req.Context(), "success, req:%+v, code:%d", reqp, resp.StatusCode)
	}
	return resp, err
}
