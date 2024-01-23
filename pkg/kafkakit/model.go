package kafkakit

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

// Msg 简单的消息格式
type Msg struct {
	Time      time.Time
	Topic     string
	Partition int
	Offset    int64
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type Model interface {
	TopicName() string
}

type FailModel interface {
	Handle(ctx context.Context, msg kafka.Message, retry int, hErr error)
}
