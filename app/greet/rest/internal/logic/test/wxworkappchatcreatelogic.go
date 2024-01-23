package test

import (
	"context"
	"fmt"
	"go-zero-yun/plugin/wxwork"
	"go-zero-yun/public/logic/pluginlgc/wxworklgc"
	"time"

	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxWorkAppChatCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxWorkAppChatCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxWorkAppChatCreateLogic {
	return &WxWorkAppChatCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxWorkAppChatCreateLogic) WxWorkAppChatCreate(req *types.TestWxWorkAppChatCreateReq) (resp *types.TestWxWorkAppChatCreateRsp, err error) {
	agent, err := wxworklgc.GetAgent(l.ctx)
	if err != nil {
		return
	}
	rs, err := agent.AppChatCreate(wxwork.AppChatCreateReq{
		Name:     req.Name,
		Owner:    req.Owner,
		Userlist: req.Userlist,
		//Userlist: []string{"MaoMaoWuXin", "ksana", "LinShengJie"},
		Chatid: "",
	})
	if err != nil {
		return
	}
	content := fmt.Sprintf("这是测试消息,当前时间:%s", time.Now().Format(time.RFC3339))
	_, err = agent.AppChatSendText(rs.Chatid, content)
	if err != nil {
		return
	}
	resp = &types.TestWxWorkAppChatCreateRsp{
		Chatid: rs.Chatid,
	}
	return
}
