package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	RequestIdKey = "X-Request-Id"
	CodeOk       = 2000
	CodeErr      = -1
)

type ReqParams struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	Header http.Header `json:"header"`
}

type TraceResponse[T any] struct {
	xhttp.BaseResponse[T]
	TraceInfo TraceInfo `json:"trace_info"`
}

type TraceInfo struct {
	Trace     string  `json:"trace"`
	Span      string  `json:"span"`
	Runtime   float64 `json:"runtime"`
	RequestId string  `json:"request_id"`
}

// Request 通用请求处理
func Request(r *http.Request, req interface{}) *http.Request {
	rp := ReqParams{
		Method: r.Method,
		Header: r.Header,
		Params: req,
	}
	fields := []logx.LogField{
		{Key: logkit.RouteName, Value: r.URL.Path},
	}
	reqId := r.Header.Get(RequestIdKey)
	if reqId != "" {
		fields = append(fields, logx.LogField{Key: logkit.RequestIdName, Value: reqId})
	}
	// 写入ctx变量
	ctx := logx.WithFields(r.Context(), fields...)
	r2 := r.WithContext(ctx)
	//记录日志,type为request
	logkit.WithType(logkit.LogRequest).Infof(ctx, logkit.LogRequest, "%+v", rp)
	return r2
}

// Response 通用返回
func Response(w http.ResponseWriter, r *http.Request, resp interface{}, err error, timer *utils.ElapsedTimer) {
	var runtime float64
	if timer != nil {
		runtime = funckit.DurToMic2(timer.Duration())
	} else {
		runtime = -1
	}
	trp := TraceResponse[any]{
		TraceInfo: TraceInfo{
			Span:      trace.SpanIDFromContext(r.Context()),
			Trace:     trace.TraceIDFromContext(r.Context()),
			Runtime:   runtime,
			RequestId: r.Header.Get(RequestIdKey),
		},
	}
	if err != nil {
		// code-data 错误响应
		trp.BaseResponse = wrapBaseResponse(err)
		httpx.OkJsonCtx(r.Context(), w, trp)
		logkit.WithType(logkit.LogResponse).WithRuntime(runtime).Errorf(r.Context(), "err: %+v, code: %d", err, trp.Code)
	} else {
		// code-data 正确响应
		trp.BaseResponse = wrapBaseResponse(resp)
		httpx.OkJsonCtx(r.Context(), w, trp)
		logkit.WithType(logkit.LogResponse).WithRuntime(runtime).Infof(r.Context(), "%+v", resp)
	}
}

// JwtFailRsp jwt验证失败
func JwtFailRsp(w http.ResponseWriter, r *http.Request, err error) {
	rp := ReqParams{
		Method: r.Method,
		Header: r.Header,
	}
	fields := []logx.LogField{
		{Key: logkit.RouteName, Value: r.URL.Path},
	}
	reqId := r.Header.Get(RequestIdKey)
	if reqId != "" {
		fields = append(fields, logx.LogField{Key: logkit.RequestIdName, Value: reqId})
	}
	// 写入ctx变量
	ctx := logx.WithFields(r.Context(), fields...)
	trp := TraceResponse[any]{
		TraceInfo: TraceInfo{
			Trace:     trace.TraceIDFromContext(r.Context()),
			Span:      trace.SpanIDFromContext(r.Context()),
			Runtime:   -1,
			RequestId: reqId,
		},
		BaseResponse: wrapBaseResponse(err),
	}
	httpx.WriteJsonCtx(ctx, w, http.StatusUnauthorized, trp)
	logkit.Errorf(ctx, "jwt fail: %v,request params: %+v", err, rp)
	return
}

// wrapBaseResponse 解析函数
func wrapBaseResponse(v any) xhttp.BaseResponse[any] {
	var resp xhttp.BaseResponse[any]
	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case *status.Status:
		resp.Code = int(data.Code())
		resp.Msg = data.Message()
	case error:
		resp.Code = CodeErr
		resp.Msg = data.Error()
	default:
		resp.Code = CodeOk
		resp.Msg = xhttp.BusinessMsgOk
		resp.Data = v
	}
	return resp
}
