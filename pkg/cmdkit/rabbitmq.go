package cmdkit

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	pconf "go-zero-yun/public/config"
)

// initRabbitmq 初始化rabbitmq
func initRabbitmq(check bool) {
	pconf.ClientCfg.Rabbitmq = make(map[string]*pconf.ClientRabbitmqConfig)
	for _, v := range pconf.Cfg.Client.Rabbitmq {
		pconf.ClientCfg.Rabbitmq[v.Name] = &pconf.ClientRabbitmqConfig{
			RabbitmqCfg: v,
		}
		if check && !IsEnvLocal() { //检查连接
			_, err := RabbitmqConnect(pconf.ClientCfg.Rabbitmq[v.Name])
			if err != nil {
				panic(errors.New(fmt.Sprintf("rabbitmq fatal %s, name: %s", err.Error(), v.Name)))
			}
		}
	}
}

// 关闭Rabbitmq
func closeRabbitmq() {
	for _, v := range pconf.ClientCfg.Rabbitmq {
		if v.Connection != nil {
			_ = v.Connection.Close()
		}
	}
}

// RabbitmqConnect 检查连接是否可用
func RabbitmqConnect(v *pconf.ClientRabbitmqConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(v.Url)
	if conn != nil {
		v.Connection = conn
	}
	return conn, err
}

// RabbitmqChannel rabbitmq通道
func RabbitmqChannel(v *pconf.ClientRabbitmqConfig) (*amqp.Channel, error) {
	if v.Connection == nil {
		conn, err := amqp.Dial(v.Url)
		if err != nil {
			return nil, err
		}
		v.Connection = conn
	}
	channel, err := v.Connection.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}
