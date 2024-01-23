package wx

import (
	"go-zero-yun/plugin"
)

// Config WX配置
type Config struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

var configMp = make(map[string]Config)

// GetCfgByName 获取配置
func GetCfgByName(name string) Config {
	if name == "" {
		name = plugin.DefaultName
	}

	if _, ok := configMp[name]; ok {
		return configMp[name]
	} else {
		return Config{}
	}
}
