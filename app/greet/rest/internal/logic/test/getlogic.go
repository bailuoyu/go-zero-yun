package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	zerrs "github.com/zeromicro/x/errors"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/public/logic/client/xormlgc"
	"go-zero-yun/public/model/db/core"
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

func (l *GetLogic) Get(req *types.TestGetReq) (resp *types.TestGetRsp, err error) {
	eg := xormlgc.Core()
	var uc core.UserAccount
	ses := eg.Context(l.ctx)
	has, err := ses.Where("id=?", req.Id).Get(&uc)
	if err != nil {
		err = zerrs.New(-5, err.Error()) //自定义code
		return
	}
	if !has {
		return
	}
	resp = &types.TestGetRsp{
		Id:       uc.Id,
		Phone:    uc.Phone,
		NickName: uc.Nickname,
	}
	return
}
