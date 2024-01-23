package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/pkg/logkit"
	"go-zero-yun/public/logic/client/elasticlgc"
	"go-zero-yun/public/model/es/core"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElasticAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewElasticAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElasticAddLogic {
	return &ElasticAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ElasticAddLogic) ElasticAdd(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	acl := core.AccessLog{
		Time:    time.Now(),
		Trace:   trace.TraceIDFromContext(l.ctx),
		Span:    trace.SpanIDFromContext(l.ctx),
		Route:   "/web/test/elastic",
		UserId:  1,
		Level:   "info",
		Type:    logkit.LogDefault,
		Runtime: 0.1,
		Content: "测试请求日志",
		Caller:  "",
	}
	esc := elasticlgc.Core()
	_, err = esc.Index().Index(acl.IndexName()).BodyJson(acl).Do(l.ctx)
	//判断版本号
	//if req.Version > 0 {
	//	ins.VersionType("external").Version(req.Version)
	//}
	return
}
