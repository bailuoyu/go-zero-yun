package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
)

// RedisCfg redis配置结构体
type RedisCfg struct {
	Name string `json:"Name"`
	redis.RedisConf
	Db int `json:"Db,default=0"`
}

// ClientRedisConfig 客户端Redis配置
type ClientRedisConfig struct {
	RedisCfg
	Client  *redis.Redis
	RwMutex sync.RWMutex
}
