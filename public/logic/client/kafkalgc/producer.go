package kafkalgc

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go-zero-yun/pkg/kafkakit"
)

// GetProducer 生产者
func GetProducer(ctx context.Context, writer *kafka.Writer) kafkakit.Producer {
	return kafkakit.Producer{
		Writer: writer,
		Ctx:    ctx,
	}
}
