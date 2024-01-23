package config

import (
	"github.com/olivere/elastic/v7"
	"sync"
)

// ElasticCfg elastic配置结构体
type ElasticCfg struct {
	Name     string   `json:"Name"`
	Urls     []string `json:"Urls"`
	Username string   `json:"Username"`
	Password string   `json:"Password"`
	Sniff    bool     `json:"Sniff"`
}

// ClientElasticConfig 客户端elastic配置
type ClientElasticConfig struct {
	ElasticCfg
	Client  *elastic.Client
	RwMutex sync.RWMutex
}
