package kafkakit

import (
	"github.com/segmentio/kafka-go"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"golang.org/x/net/context"
)

// Producer 生产者
type Producer struct {
	Writer *kafka.Writer
	Ctx    context.Context
}

// Close 关闭生产者
func (prd *Producer) Close() error {
	return prd.Writer.Close()
}

// Push 推送消息
func (prd *Producer) Push(msgs ...kafka.Message) error {
	if prd.Writer.Completion == nil {
		prd.Writer.Completion = prd.writerCompletion
	}
	// 添加MessageId
	for i := range msgs {
		msgs[i].Headers = append(msgs[i].Headers, kafka.Header{
			Key:   MessageIdKey,
			Value: []byte(funckit.RandomTimeStr(12, 1)),
		})
	}
	return prd.Writer.WriteMessages(prd.Ctx, msgs...)
}

// writerCompletion 注入函数，主要用于写日志
func (prd *Producer) writerCompletion(messages []kafka.Message, err error) {
	// 日志
	//var msgs []Msg
	//for _, v := range messages {
	//	msgs = append(msgs, Msg{
	//		Key:   string(v.Key),
	//		Value: string(v.Value),
	//	})
	//}
	if err != nil {
		logkit.WithType(logkit.LogKafkaWrite).Errorf(prd.Ctx, "write message error:%v", err)
		return
	}
	logkit.WithType(logkit.LogKafkaWrite).Info(prd.Ctx, "write message success")
}
