package config

import (
    conf "go-zero-yun/public/config"
    {{.authImport}}
)

var (
	CfgFile  string
	c        Config
	LoadOnce sync.Once
)

type Config struct {
	conf.Config
	Custom struct{} //自定义配置
	{{.auth}}
	{{.jwtTrans}}
}

// LoadCfg 载入配置
func LoadCfg() Config {
	var c Config
	conf.MustLoad(CfgFile, &c)
	return c
}

// LoadCfgOnce 防止协程中多次加载
func LoadCfgOnce() {
	// 防止协程中多次加载
	LoadOnce.Do(func() {
		LoadCfg()
	})
}

// GetCfg 获取配置
func GetCfg() Config {
	return c
}