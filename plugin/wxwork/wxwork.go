// Package wxwork 企业微信包
package wxwork

import (
	"context"
	"github.com/imroc/req/v3"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-yun/pkg/reqkit"
	"go-zero-yun/plugin"
)

type Agent struct {
	Config Config
	RdsCli *redis.Redis
	Token  Token
	ReqC   *req.Client
	ctx    context.Context
}

func GetAgent(ctx context.Context, rdsCli *redis.Redis) (Agent, error) {
	return GetAgentByName(ctx, rdsCli, plugin.DefaultName)
}

func GetAgentByName(ctx context.Context, rdsCli *redis.Redis, name string) (Agent, error) {
	agent := Agent{
		Config: GetCfgByName(name),
		RdsCli: rdsCli,
		ReqC:   reqkit.Client(),
		ctx:    ctx,
	}
	err := agent.initToken()
	return agent, err
}

func (agent *Agent) GetContext() context.Context {
	return agent.ctx
}
