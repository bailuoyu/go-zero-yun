package wxwork

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	redisKeyPre  = "wx_work:agent"
	redisLockPre = "wx_work:lock"
)

type Token struct {
	AccessToken string
	ExpiresIn   int
}

// InitToken 初始化token
func (agent *Agent) initToken() error {
	err := agent.getTokenFromRedis()
	if err == redis.Nil || agent.Token.ExpiresIn <= 10 {
		err = agent.setToken()
	}
	return err
}

// getTokenFromRedis 从redis获取token
func (agent *Agent) getTokenFromRedis() error {
	redisKey := agent.getRedisKey()
	act, err := agent.RdsCli.GetCtx(agent.ctx, redisKey)
	var ex int
	if err == nil {
		ex, err = agent.RdsCli.TtlCtx(agent.ctx, redisKey)
		agent.Token = Token{
			AccessToken: act,
			ExpiresIn:   ex,
		}
	}
	return err
}

// getRedisKey 获取redis key
func (agent *Agent) getRedisKey() string {
	return fmt.Sprintf("%s:%s-%d", redisKeyPre, agent.Config.CorpId, agent.Config.AgentId)
}

// getLockName 获取redis lock name
func (agent *Agent) getLockName() string {
	return fmt.Sprintf("%s:%s-%d", redisLockPre, agent.Config.CorpId, agent.Config.AgentId)
}

// setToken 设置token
func (agent *Agent) setToken() error {
	// 锁名称
	lockName := agent.getLockName()
	// 加锁
	lock := redis.NewRedisLock(agent.RdsCli, lockName)
	lock.SetExpire(10)
	// 尝试获取锁
	acquire, err := lock.AcquireCtx(agent.ctx)
	switch {
	case err != nil:
		// 错误直接返回
		return err
	case acquire:
		// 获取到锁
		defer lock.Release() // 释放锁
		// 业务逻辑
	case !acquire:
		// 没有拿到锁直接请求
		return nil
	}
	redisKey := agent.getRedisKey()
	//否则请求
	rs, err := agent.GetToken()
	if err != nil {
		return err
	}
	agent.Token = Token{
		AccessToken: rs.AccessToken,
		ExpiresIn:   rs.ExpiresIn - 5,
	}
	err = agent.RdsCli.SetexCtx(agent.ctx, redisKey, agent.Token.AccessToken, agent.Token.ExpiresIn)
	return err
}

type GetTokenRsp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetToken 获取access_token https://developer.work.weixin.qq.com/document/path/91039
func (agent *Agent) GetToken() (GetTokenRsp, error) {
	url := agent.getUrl("gettoken")
	request := agent.ReqC.Get(url)
	var rs GetTokenRsp
	response := request.SetQueryParams(map[string]string{
		"corpid":     agent.Config.CorpId,
		"corpsecret": agent.Config.Secret,
	}).Do()
	err := agent.processRsp(request, response)
	if err != nil {
		return rs, err
	}
	err = response.UnmarshalJson(&rs)
	return rs, err
}
