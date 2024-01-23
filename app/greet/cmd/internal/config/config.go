package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	pconf "go-zero-yun/public/config"
	"sync"
)

var (
	CfgFile  string
	c        Config
	LoadOnce sync.Once
)

type Config struct {
	pconf.Config
	Custom struct{} //自定义配置
}

// LoadCfg 载入配置
func LoadCfg() Config {
	conf.MustLoad(CfgFile, &c)
	// 全局配置
	pconf.Cfg = &c.Config
	// 加载日志配置
	logx.MustSetup(c.Config.Server.Cmd.Log)
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
