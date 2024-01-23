package monlgc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetDatabase @Description
func GetDatabase(name string) *mongo.Database {
	if _, ok := conf.ClientCfg.Mongo[name]; !ok {
		logx.Errorw(fmt.Sprintf("mongo empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogMongo})
		return &mongo.Database{}
	}
	if conf.ClientCfg.Mongo[name].Database != nil {
		return conf.ClientCfg.Mongo[name].Database
	}
	// 如果没有连接
	// 加锁防止重复建立连接
	conf.ClientCfg.Mongo[name].RwMutex.Lock()
	defer conf.ClientCfg.Mongo[name].RwMutex.Unlock()
	if conf.ClientCfg.Mongo[name].Database == nil {
		_, err := cmdkit.MongoConnect(conf.ClientCfg.Mongo[name])
		if err != nil {
			logx.Errorw(fmt.Sprintf("mongo fatal %s, name: %s", err.Error(), name),
				logx.LogField{Key: logkit.TypeName, Value: logkit.LogMongo})
		}
	}
	return conf.ClientCfg.Mongo[name].Database
}

// Core 获取核心nosql
func Core() *mongo.Database {
	return GetDatabase("core")
}
