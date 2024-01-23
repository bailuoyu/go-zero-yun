package dy

import (
	"go-zero-yun/plugin"
)

// Config DY配置
type Config struct {
	ClientSecret string `json:"client_secret"`
	ClientKey    string `json:"client_key"`
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
