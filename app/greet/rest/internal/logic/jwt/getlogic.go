package jwt

import (
	"context"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
