package qq

import (
	"go-zero-yun/plugin"
)

// Config QQ配置
type Config struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
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
