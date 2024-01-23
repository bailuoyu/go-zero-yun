package config

import (
	"sync"
	"xorm.io/xorm"
)

// MysqlCfg mysql配置结构体
type MysqlCfg struct {
	Name         string   `json:"Name"`
	Dsn          string   `json:"Dsn"`
	SlaveDsns    []string `json:"SlaveDsns,optional"`
	MaxIdleConns int      `json:"MaxIdleConns"`
	MaxOpenConns int      `json:"MaxOpenConns"`
	ShowSql      bool     `json:"ShowSql,optional"`
	LogLevel     string   `json:"LogLevel,optional"`
}

// ClientMysqlConfig 客户端mysql配置
type ClientMysqlConfig struct {
	MysqlCfg
	Engine  *xorm.EngineGroup
	RwMutex sync.RWMutex
}
