package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type RabbitmqCfg struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

// ClientRabbitmqConfig 客户端rabbitmq配置
type ClientRabbitmqConfig struct {
	RabbitmqCfg
	Connection *amqp.Connection
	Channel    *amqp.Channel
	RWMutex    sync.RWMutex
}
