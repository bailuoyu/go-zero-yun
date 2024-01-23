package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-yun/plugin"
)

var Cfg *Config
var ClientCfg ClientConfig

type Config struct {
	Global Global `json:"Global"`
	Server struct {
		App  string             `json:"App"`
		Rest rest.RestConf      `json:"Rest,optional"`
		Zrpc zrpc.RpcServerConf `json:"Zrpc,optional"`
		Cmd  CmdConfig          `json:"Cmd,optional"`
	}
	Client struct {
		Mysql    []MysqlCfg    `json:"Mysql,optional"`
		Redis    []RedisCfg    `json:"Redis,optional"`
		Mongo    []MongoCfg    `json:"Mongo,optional"`
		Kafka    []KafkaCfg    `json:"Kafka,optional"`
		Rabbitmq []RabbitmqCfg `json:"Rabbitmq,optional"`
		Elastic  []ElasticCfg  `json:"Elastic,optional"`
	} `json:"Client,optional"`
	Pkg    PkgCfg        `json:"Pkg,optional"`
	Plugin plugin.Config `json:"Plugin,optional"`
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Mysql    map[string]*ClientMysqlConfig
	Redis    map[string]*ClientRedisConfig
	Mongo    map[string]*ClientMongoConfig
	Kafka    map[string]KafkaCfg
	Rabbitmq map[string]RabbitmqCfg
	Elastic  map[string]*ClientElasticConfig
}
