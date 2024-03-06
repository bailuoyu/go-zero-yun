package rabbitmqlgc

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-zero-yun/pkg/rabbitmqkit"
)

// GetProducer 获取生产者
func GetProducer(ctx context.Context, channel *amqp.Channel, model interface{}) rabbitmqkit.Producer {
	producer := rabbitmqkit.Producer{
		Channel: channel,
		Ctx:     ctx,
	}
	if model == nil {
		return producer
	}
	switch mt := model.(type) {
	case rabbitmqkit.QueueModel:
		producer.RoutingKey = mt.QueueName()
	case rabbitmqkit.ExchangeModel:
		producer.Exchange = mt.ExchangeName()
	case rabbitmqkit.RoutingKeyModel:
		producer.Exchange = mt.Exchange().ExchangeName()
		producer.RoutingKey = mt.RoutingKey()
	}
	return producer
}
