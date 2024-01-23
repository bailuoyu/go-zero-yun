package redislgc

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
)

// GetClient 获取redis引擎
func GetClient(name string) *redis.Redis {

	if _, ok := conf.ClientCfg.Redis[name]; !ok {
		logx.Errorw(fmt.Sprintf("redis empty, name: %s, cfg: %v", name, conf.ClientCfg.Redis),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogRedis})
		return &redis.Redis{}
	}
	if conf.ClientCfg.Redis[name].Client != nil {
		return conf.ClientCfg.Redis[name].Client
	}
	// 如果没有连接
	// 加锁防止重复建立连接
	conf.ClientCfg.Redis[name].RwMutex.Lock()
	defer conf.ClientCfg.Redis[name].RwMutex.Unlock()
	if conf.ClientCfg.Redis[name].Client == nil {
		_, ok := cmdkit.RedisConnect(conf.ClientCfg.Redis[name], true)
		if !ok {
			logx.Errorw(fmt.Sprintf("redis fatal, name: %s", name),
				logx.LogField{Key: logkit.TypeName, Value: logkit.LogRedis})
		}
	}
	return conf.ClientCfg.Redis[name].Client
}

// Core 获取核心nosql
func Core() *redis.Redis {
	return GetClient("core")
}

// Data 获取数据nosql
func Data() *redis.Redis {
	return GetClient("data")
}
