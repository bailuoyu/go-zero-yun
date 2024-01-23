package rabbitmqlgc

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	"go-zero-yun/pkg/rabbitmqkit"
	conf "go-zero-yun/public/config"
)

// GetChannel 获取rabbitmq通道
func GetChannel(name string) rabbitmqkit.Client {
	client := rabbitmqkit.Client{
		Con:     &amqp.Connection{},
		Channel: &amqp.Channel{},
	}
	if _, ok := conf.ClientCfg.Rabbitmq[name]; !ok {
		logx.Errorw(fmt.Sprintf("rabbitmq empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogRabbitmq},
		)
		return client
	}
	var err error
	client.Con, client.Channel, err = cmdkit.RabbitmqChannel(conf.ClientCfg.Rabbitmq[name])
	if err != nil {
		logx.Errorw(fmt.Sprintf("rabbitmq fatal %s, name: %s", err.Error(), name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogRabbitmq},
		)
	}
	return client
}

// Core 获取核心连接通道
func Core() rabbitmqkit.Client {
	return GetChannel("core")
}
