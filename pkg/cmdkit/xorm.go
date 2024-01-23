package cmdkit

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-zero-yun/pkg/logkit"
	pconf "go-zero-yun/public/config"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

// XormLogLevelMp 日志等级map
var XormLogLevelMp = map[string]log.LogLevel{
	"debug":   log.LOG_DEBUG,
	"info":    log.LOG_INFO,
	"warning": log.LOG_WARNING,
	"error":   log.LOG_ERR,
	"off":     log.LOG_OFF,
}

// initXorm 初始化xorm
func initXorm(check bool) {
	pconf.ClientCfg.Mysql = make(map[string]*pconf.ClientMysqlConfig)
	for _, v := range pconf.Cfg.Client.Mysql {
		pconf.ClientCfg.Mysql[v.Name] = &pconf.ClientMysqlConfig{
			MysqlCfg: v,
		}
		_, err := XormConnect(pconf.ClientCfg.Mysql[v.Name], check)
		if err != nil {
			panic(errors.New(fmt.Sprintf("xorm fatal %s, name: %s", err.Error(), v.Name)))
		}
	}
}

// closeXorm 关闭Xorm
func closeXorm() {
	for _, v := range pconf.ClientCfg.Mysql {
		if v != nil {
			v.Engine.Close()
		}
	}
}

// XormConnect 获取xorm连接
func XormConnect(v *pconf.ClientMysqlConfig, check bool) (*xorm.EngineGroup, error) {
	dsns := []string{v.Dsn}
	if len(v.SlaveDsns) > 0 {
		dsns = append(dsns, v.SlaveDsns...)
	}
	engineGroup, err := xorm.NewEngineGroup("mysql", dsns)
	if err != nil {
		return nil, err
	}
	//最大连接数
	if v.MaxIdleConns > 0 {
		engineGroup.SetMaxIdleConns(v.MaxIdleConns)
	}
	if v.MaxOpenConns > 0 {
		engineGroup.SetMaxOpenConns(v.MaxOpenConns)
	}
	engineGroup.ShowSQL(v.ShowSql)
	engineGroup.Logger().SetLevel(GetLogLevel(v.LogLevel))
	//设置日志
	logkit.SetXormGroupLog(engineGroup)
	v.Engine = engineGroup
	// ping一次连接确保能连接上
	if !check || IsEnvLocal() {
		return engineGroup, err
	}
	err = engineGroup.Ping()
	return engineGroup, err
}

// GetLogLevel 获取日志等级
func GetLogLevel(logLevel string) log.LogLevel {
	if v, ok := XormLogLevelMp[logLevel]; ok {
		return v
	}
	return log.LOG_INFO
}
