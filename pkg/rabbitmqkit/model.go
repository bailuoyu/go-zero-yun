package rabbitmqkit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueModel interface {
	QueueName() string
}

type ExchangeModel interface {
	ExchangeName() string
}

type RoutingKeyModel interface {
	RoutingKey() string
	Exchange() ExchangeModel
}

type FailModel interface {
	Handle(ctx context.Context, msg amqp.Delivery, retry int, hErr error)
}
