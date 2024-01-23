package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/pkg/rediskit"
	"go-zero-yun/public/logic/client/redislgc"
	"go-zero-yun/public/model/redis/core"
)

type RedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisLogic {
	return &RedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisLogic) Redis(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	redis := redislgc.Core()
	key, sec := rediskit.GetInfo(core.UserInfo{}, "1")
	err = redis.SetexCtx(l.ctx, key, "abc", sec)
	return
}
