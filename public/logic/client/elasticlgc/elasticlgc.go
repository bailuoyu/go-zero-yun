package elasticlgc

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
)

// GetClient 获取Elastic引擎
func GetClient(name string) *elastic.Client {
	if _, ok := conf.ClientCfg.Elastic[name]; !ok {
		logx.Errorw(fmt.Sprintf("elastic empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogES})
		return &elastic.Client{}
	}
	if conf.ClientCfg.Elastic[name].Client != nil {
		return conf.ClientCfg.Elastic[name].Client
	}
	// 如果没有连接
	// 加锁防止重复建立连接
	conf.ClientCfg.Elastic[name].RwMutex.Lock()
	defer conf.ClientCfg.Elastic[name].RwMutex.Unlock()
	if conf.ClientCfg.Elastic[name].Client == nil {
		_, err := cmdkit.ElasticConnect(conf.ClientCfg.Elastic[name])
		if err != nil {
			logx.Errorw(fmt.Sprintf("elastic fatal, name: %s", name),
				logx.LogField{Key: logkit.TypeName, Value: logkit.LogES})
		}
	}
	return conf.ClientCfg.Elastic[name].Client
}

// Core 获取核心nosql
func Core() *elastic.Client {
	return GetClient("core")
}
