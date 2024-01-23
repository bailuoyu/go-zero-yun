package test

import (
	"context"
	"fmt"
	"go-zero-yun/public/logic/pluginlgc/wxworklgc"
	"time"

	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxWorkAppChatSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxWorkAppChatSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxWorkAppChatSendLogic {
	return &WxWorkAppChatSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxWorkAppChatSendLogic) WxWorkAppChatSend(req *types.Empty) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	agent, err := wxworklgc.GetAgent(l.ctx)
	if err != nil {
		return
	}
	chatid := wxworklgc.GetChatWarningId()
	content := fmt.Sprintf("这是测试消息,当前时间:%s", time.Now().Format(time.RFC3339))
	_, err = agent.AppChatSendText(chatid, content)
	if err != nil {
		return
	}
	return
}
