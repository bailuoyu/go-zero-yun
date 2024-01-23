package cmdkit

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	pconf "go-zero-yun/public/config"
)

// initRabbitmq 初始化rabbitmq
func initRabbitmq(check bool) {
	pconf.ClientCfg.Rabbitmq = make(map[string]pconf.RabbitmqCfg)
	for _, v := range pconf.Cfg.Client.Rabbitmq {
		pconf.ClientCfg.Rabbitmq[v.Name] = v
		if check && !IsEnvLocal() { //检查连接
			err := RabbitmqConnect(v)
			if err != nil {
				panic(errors.New(fmt.Sprintf("rabbitmq fatal %s, name: %s", err.Error(), v.Name)))
			}
		}
	}
}

// RabbitmqConnect 检查连接是否可用
func RabbitmqConnect(v pconf.RabbitmqCfg) error {
	conn, err := amqp.Dial(v.Url)
	if conn != nil {
		defer conn.Close()
	}
	return err
}

// RabbitmqChannel rabbitmq通道
func RabbitmqChannel(v pconf.RabbitmqCfg) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(v.Url)
	if err != nil {
		return nil, nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return conn, nil, err
	}
	return conn, channel, nil
}
