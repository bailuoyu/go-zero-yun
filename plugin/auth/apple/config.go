package apple

import (
	"go-zero-yun/plugin"
)

// Config APPLE配置
type Config struct {
	KeysUrl string `json:"keys_url"`
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
