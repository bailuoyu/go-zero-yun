package kafkalgc

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/kafkakit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
)

// GetWriter 获取Kafka引擎
func GetWriter(name string, model kafkakit.Model) *kafka.Writer {
	if _, ok := conf.ClientCfg.Kafka[name]; !ok {
		logx.Errorw(fmt.Sprintf("kafka empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogKafkaWrite},
		)
		return &kafka.Writer{}
	}
	var topic string
	if model != nil {
		topic = model.TopicName()
	}
	return cmdkit.KafkaWriter(conf.ClientCfg.Kafka[name], topic)
}

// GetReader 获取Kafka引擎
func GetReader(name string, models ...kafkakit.Model) *kafka.Reader {
	if _, ok := conf.ClientCfg.Kafka[name]; !ok {
		logx.Errorw(fmt.Sprintf("kafka empty, name: %s", name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogKafkaRead},
		)
		return &kafka.Reader{}
	}
	var topics []string
	if len(models) == 0 {
		return &kafka.Reader{}
	}
	for _, model := range models {
		topics = append(topics, model.TopicName())
	}
	return cmdkit.KafkaReader(conf.ClientCfg.Kafka[name], topics...)
}

// CoreWriter 获取core的Kafka writer
func CoreWriter(model kafkakit.Model) *kafka.Writer {
	return GetWriter("core", model)
}

// CoreReader 获取core的Kafka reader
func CoreReader(models ...kafkakit.Model) *kafka.Reader {
	return GetReader("core", models...)
}
