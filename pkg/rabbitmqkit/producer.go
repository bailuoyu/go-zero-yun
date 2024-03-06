package rabbitmqkit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"time"
)

// Producer 生产者
type Producer struct {
	Channel    *amqp.Channel
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Ctx        context.Context
}

// Close 关闭生产者,生产者无需关闭
//func (prd *Producer) Close() error {
//	return prd.Channel.Close()
//}

// Push 推送消息
func (prd *Producer) Push(msg amqp.Publishing) error {
	if msg.ContentType == "" {
		msg.ContentType = defaultContentType
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}
	if msg.MessageId == "" {
		msg.MessageId = funckit.RandomTimeStr(12, 1)
	}
	if msg.DeliveryMode == 0 {
		msg.DeliveryMode = amqp.Persistent
	}
	err := prd.Channel.PublishWithContext(
		prd.Ctx,
		prd.Exchange,
		prd.RoutingKey,
		false,
		false,
		msg,
	)
	if err == nil {
		logkit.WithType(logkit.LogRabbitmqWrite).Info(prd.Ctx, "write message success")
	} else {
		logkit.WithType(logkit.LogRabbitmqWrite).Errorf(prd.Ctx, "write message error:%v", err)
	}
	return err
}

// DelayPush 延时消息
func (prd *Producer) DelayPush(msg amqp.Publishing, duration time.Duration) error {
	DelayMsg(&msg, duration)
	return prd.Push(msg)
}
