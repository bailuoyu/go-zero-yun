package rabbitmqkit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

const (
	defaultContentType = "text/plain"
	XDelayKey          = "x-delay"
)

type Client struct {
	Con     *amqp.Connection
	Channel *amqp.Channel
}

// Msg 简单的消息格式
type Msg struct {
	Time       time.Time
	Exchange   string
	RoutingKey string
	MessageId  string
	Body       string
}

// DelayMsg 添加延时
func DelayMsg(msg *amqp.Publishing, duration time.Duration) {
	if msg.Headers == nil {
		msg.Headers = make(amqp.Table)
	}
	msg.Headers[XDelayKey] = int(duration.Milliseconds())
}
