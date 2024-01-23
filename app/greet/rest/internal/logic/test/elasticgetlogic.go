package test

import (
	"context"
	"encoding/json"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/public/logic/client/elasticlgc"
	"go-zero-yun/public/model/es/core"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElasticGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewElasticGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElasticGetLogic {
	return &ElasticGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ElasticGetLogic) ElasticGet(req *types.TestElasticGetReq) (resp *types.TestElasticGetRsp, err error) {
	// todo: add your logic here and delete this line
	id := req.Id
	var acl core.AccessLog
	esc := elasticlgc.Core()
	res, err := esc.Get().Index(acl.IndexName()).Id(id).Do(l.ctx)
	if err != nil {
		return
	}
	if err = json.Unmarshal(res.Source, &acl); err != nil {
		return
	}
	resp = &types.TestElasticGetRsp{
		Time:    acl.Time.Format(time.RFC3339),
		Trace:   acl.Trace,
		Span:    acl.Span,
		Route:   acl.Route,
		UserId:  acl.UserId,
		Level:   acl.Level,
		Type:    acl.Type,
		Runtime: acl.Runtime,
		Content: acl.Content,
		Caller:  acl.Caller,
	}
	return
}

func (l *ElasticGetLogic) demo() {
	esc := elasticlgc.Core()
	//json查询语句，json写在代码里可读性很差，用map代替
	source := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"search_name": "catcc",
			},
		},
		"size":    20,
		"version": true,
	}
	searchResult, err := esc.Search().Index("test").Source(source).Size(1).Do(l.ctx)
	if err != nil {
		return
	}
	// 如果查询结果为空
	if searchResult.Hits == nil {
		return
	}
	for _, hit := range searchResult.Hits.Hits {
		println(hit)
	}
}
