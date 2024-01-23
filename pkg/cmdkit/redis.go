package cmdkit

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-yun/pkg/logkit"
	pconf "go-zero-yun/public/config"
)

// initRedis 初始化redis
func initRedis(check bool) {
	pconf.ClientCfg.Redis = make(map[string]*pconf.ClientRedisConfig)
	for _, v := range pconf.Cfg.Client.Redis {
		pconf.ClientCfg.Redis[v.Name] = &pconf.ClientRedisConfig{
			RedisCfg: v,
		}
		_, ok := RedisConnect(pconf.ClientCfg.Redis[v.Name], check)
		if !ok {
			panic(errors.New(fmt.Sprintf("redis fatal ping.api, name: %s", v.Name)))
		}
	}
}

// closeRedis 关闭Redis
func closeRedis() {
	for _, v := range pconf.ClientCfg.Redis {
		if v != nil {
			v.Client.Pipelined(func(p redis.Pipeliner) error {
				p.Close()
				return nil
			})
		}
	}
}

// RedisConnect 获取redis连接
func RedisConnect(v *pconf.ClientRedisConfig, check bool) (*redis.Redis, bool) {
	client := redis.MustNewRedis(v.RedisConf)
	v.Client = client
	if v.Db != 0 {
		v.Client.Pipelined(func(p redis.Pipeliner) error {
			p.Select(context.Background(), v.Db)
			return nil
		})
	}
	//设置redis日志
	logkit.SetRedisLog()
	// ping一次连接确保能连接上
	if !check || IsEnvLocal() {
		return client, true
	}
	ok := client.Ping()
	return client, ok
}
