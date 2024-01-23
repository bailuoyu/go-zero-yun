package test

import (
	"context"
	"fmt"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/plugin/cos"

	"github.com/zeromicro/go-zero/core/logx"
)

type CosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CosLogic {
	return &CosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CosLogic) Cos(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	CosC := cos.GetClient(l.ctx)
	fmt.Println(CosC.Config)
	return
}
