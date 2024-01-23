package rabbitmqlgc

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-zero-yun/pkg/rabbitmqkit"
	"go-zero-yun/public/logic/client/xormlgc"
)

// GetConsumer 获取消费者
func GetConsumer(ctx context.Context, client rabbitmqkit.Client, model rabbitmqkit.QueueModel,
	handle func(ctx context.Context, msg amqp.Delivery) error) rabbitmqkit.Consumer {
	return rabbitmqkit.Consumer{
		Client:    client,
		Handle:    handle,
		RetryMax:  rabbitmqkit.DefaultRetryMax,
		Queue:     model.QueueName(),
		FailModel: rabbitmqkit.XormFailModel{Engine: xormlgc.Data().Master()},
		Ctx:       ctx,
	}
}
