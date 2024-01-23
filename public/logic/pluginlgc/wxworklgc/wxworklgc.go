package wxworklgc

import (
	"context"
	"go-zero-yun/plugin/wxwork"
	"go-zero-yun/public/config"
	"go-zero-yun/public/logic/client/redislgc"
)

// GetAgent 获取企业微信应用
func GetAgent(ctx context.Context) (wxwork.Agent, error) {
	rdsCli := redislgc.Data()
	return wxwork.GetAgent(ctx, rdsCli)
}

// GetChatWarningId 告警企业微信群id
func GetChatWarningId() string {
	return config.Cfg.Pkg.WxWork.AppChat.WarningId
}

// ChatWarning 企业微信群告警消息
func ChatWarning(ctx context.Context, content string) error {
	agent, err := GetAgent(ctx)
	if err != nil {
		return err
	}
	chatid := GetChatWarningId()
	_, err = agent.AppChatSendText(chatid, content)
	return err
}
