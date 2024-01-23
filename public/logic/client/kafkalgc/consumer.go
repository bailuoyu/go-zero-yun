package kafkalgc

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go-zero-yun/pkg/kafkakit"
	"go-zero-yun/public/logic/client/xormlgc"
)

// GetConsumer 获取消费者
func GetConsumer(ctx context.Context, reader *kafka.Reader,
	handle func(ctx context.Context, msg kafka.Message) error) kafkakit.Consumer {
	return kafkakit.Consumer{
		Reader:    reader,
		Handle:    handle,
		RetryMax:  kafkakit.DefaultRetryMax,
		FailModel: kafkakit.XormFailModel{Engine: xormlgc.Data().Master()},
		Ctx:       ctx,
	}
}
