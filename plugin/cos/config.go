package cos

import "go-zero-yun/plugin"

// Config 企业微信配置结构体
type Config struct {
	AppId     uint   `json:"AppId"`
	SecretId  string `json:"SecretId"`
	SecretKey string `json:"SecretKey"`
	Region    string `json:"Region"`
	Bucket    string `json:"Bucket"`
	//CdnDomain string `json:"CdnDomain" mapstructure:"cdn_domain"`
	CdnDomain string `json:"CdnDomain"`
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
