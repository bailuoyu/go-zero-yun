package cmdkit

import pconf "go-zero-yun/public/config"

// InitClient 初始化客户端配置
func InitClient(check bool) {
	// 初始化mysql
	if len(pconf.Cfg.Client.Mysql) > 0 {
		initXorm(check)
	}
	// 初始化redis
	if len(pconf.Cfg.Client.Redis) > 0 {
		initRedis(check)
	}
	// 初始化mongo
	if len(pconf.Cfg.Client.Mongo) > 0 {
		initMongo(check)
	}
	// 初始化kafka
	if len(pconf.Cfg.Client.Kafka) > 0 {
		initKafka(check)
	}
	// 初始化rabbitmq
	if len(pconf.Cfg.Client.Rabbitmq) > 0 {
		initRabbitmq(check)
	}
	// 初始化elastic
	if len(pconf.Cfg.Client.Elastic) > 0 {
		initElastic(check)
	}
}

// ClientConf 获取Client配置
func ClientConf() pconf.ClientConfig {
	return pconf.ClientCfg
}

// CloseClient 关闭各个客户端的长连接
func CloseClient() {
	closeXorm()
	closeRedis()
	closeMongo()
	closeKafka()
	closeRabbitmq()
}
