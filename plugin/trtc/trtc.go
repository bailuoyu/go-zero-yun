package trtc

import (
	"context"

	"github.com/imroc/req/v3"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"go-zero-yun/pkg/reqkit"
	"go-zero-yun/plugin"
)

const defaultExpire = 86400 * 180 //过期时间180天

// ClientIM TRTC客户端
type ClientTRTC struct {
	Config Config
	Sign   string
	ReqC   *req.Client
	ctx    context.Context
}

// GetClient 获取Client
func GetClient(ctx context.Context) (*ClientTRTC, error) {
	return GetClientByName(ctx, plugin.DefaultName)
}

func GetUserSign(userId string) (string, error) {
	cfg := GetCfgByName(plugin.DefaultName)
	return tencentyun.GenUserSig(cfg.SdkAppId, cfg.Key, userId, defaultExpire)
}

// GetClientByName 根据名称获取Client
func GetClientByName(ctx context.Context, name string) (*ClientTRTC, error) {
	cfg := GetCfgByName(name)
	sign, err := tencentyun.GenUserSig(cfg.SdkAppId, cfg.Key, cfg.AdminId, defaultExpire)
	if err != nil {
		return nil, err
	}
	clm := &ClientTRTC{
		Config: cfg,
		Sign:   sign,
		ReqC:   reqkit.Client(),
		ctx:    ctx,
	}
	return clm, nil
}
