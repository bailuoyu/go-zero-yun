package xormlgc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
	"xorm.io/xorm"
)

// GetEngine 获取xorm引擎
func GetEngine(name string) *xorm.EngineGroup {
	if _, ok := conf.ClientCfg.Mysql[name]; !ok {
		logx.Errorw(fmt.Sprintf("xorm empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogXorm})
		return &xorm.EngineGroup{}
	}
	if conf.ClientCfg.Mysql[name].Engine != nil {
		return conf.ClientCfg.Mysql[name].Engine
	}
	// 如果没有连接
	// 加锁防止重复建立连接
	conf.ClientCfg.Mysql[name].RwMutex.Lock()
	defer conf.ClientCfg.Mysql[name].RwMutex.Unlock()
	if conf.ClientCfg.Mysql[name].Engine == nil {
		_, err := cmdkit.XormConnect(conf.ClientCfg.Mysql[name], true)
		if err != nil {
			logx.Errorw(fmt.Sprintf("xorm fatal %s, name: %s", err.Error(), name),
				logx.LogField{Key: logkit.TypeName, Value: logkit.LogXorm})
		}
	}
	return conf.ClientCfg.Mysql[name].Engine
}

// Core 获取核心数据库
func Core() *xorm.EngineGroup {
	return GetEngine("core")
}

// Data 获取记录数据库
func Data() *xorm.EngineGroup {
	return GetEngine("data")
}
