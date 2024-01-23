package wxwork

import "go-zero-yun/plugin"

// Config 企业微信配置结构体
type Config struct {
	CorpId  string `json:"CorpId"`
	AgentId int    `json:"AgentId"`
	Secret  string `json:"Secret"`
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
