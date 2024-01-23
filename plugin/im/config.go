package im

import "go-zero-yun/plugin"

// Config IM配置
type Config struct {
	SdkAppId  int    `json:"SdkAppId"`
	AdminId   string `json:"AdminId"`
	Key       string `json:"Key"`
	ApiDomain string `json:"ApiDomain"`
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
