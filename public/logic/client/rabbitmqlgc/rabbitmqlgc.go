package rabbitmqlgc

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
)

// GetChannel 获取rabbitmq通道
func GetChannel(name string, multiplex bool) *amqp.Channel {
	if _, ok := conf.ClientCfg.Rabbitmq[name]; !ok {
		logx.Errorw(fmt.Sprintf("rabbitmq empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogRabbitmq},
		)
		return &amqp.Channel{}
	}
	if multiplex && conf.ClientCfg.Rabbitmq[name].Channel != nil && !conf.ClientCfg.Rabbitmq[name].Channel.IsClosed() {
		return conf.ClientCfg.Rabbitmq[name].Channel
	}
	channel, err := cmdkit.RabbitmqChannel(conf.ClientCfg.Rabbitmq[name])
	if err != nil {
		logx.Errorw(fmt.Sprintf("rabbitmq fatal %s, name: %s", err.Error(), name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogRabbitmq},
		)
	} else if multiplex {
		conf.ClientCfg.Rabbitmq[name].Channel = channel
	}
	return channel
}

// Core 获取核心连接通道
func Core(multiplex bool) *amqp.Channel {
	return GetChannel("core", multiplex)
}
