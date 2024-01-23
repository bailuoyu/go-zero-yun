package test

import (
	"context"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaLogic {
	return &KafkaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaLogic) Kafka(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
