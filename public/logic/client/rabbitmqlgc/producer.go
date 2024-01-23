package rabbitmqlgc

import (
	"context"
	"go-zero-yun/pkg/rabbitmqkit"
)

// GetProducer 获取生产者
func GetProducer(ctx context.Context, client rabbitmqkit.Client, model interface{}) rabbitmqkit.Producer {
	producer := rabbitmqkit.Producer{
		Client: client,
		Ctx:    ctx,
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
