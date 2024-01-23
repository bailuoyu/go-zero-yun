package test

import (
	"context"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/public/logic/client/xormlgc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.TestAddReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	eg := xormlgc.Core()
	ses := eg.NewSession().Context(l.ctx)
	ses.Begin()
	/**
	 * 业务代码
	 */
	if err != nil {
		ses.Rollback() //回滚
	} else {
		ses.Commit() //提交
	}
	return
}
