package kafkakit

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"time"
)

// BatchConsumer 批量消费者,未完成
type BatchConsumer struct {
	Reader        *kafka.Reader
	Handle        func(ctx context.Context, msgs []kafka.Message) error // bool为标明是否需要重试
	RetryMax      int                                                   // 最大重试次数
	RetryInterval time.Duration                                         // 重试间隔
	FailModel     FailModel                                             //消费失败处理函数
	Ctx           context.Context
}

// Close 关闭生产者
func (cns *BatchConsumer) Close() {
	_ = cns.Reader.Close()
}

// Run 运行消费者
func (cns *BatchConsumer) Run() error {
	//检查并设置配置默认参数
	//if cns.RetryMax == 0 {
	//	cns.RetryMax = DefaultRetryMax
	//}
	if cns.RetryMax > 0 && cns.RetryInterval == 0 {
		cns.RetryInterval = DefaultRetryInterval
	}
	// 判断是否需要生成trace
	var trace string
	if cns.Ctx == nil {
		cns.Ctx = context.Background()
	} else {
		trace = cns.Ctx.Value(logkit.TraceName).(string)
	}
	if trace == "" {
		trace = funckit.RandomStr(24, 1)
	}
	// 结束后关闭消费者
	defer cns.Close()
	for {
		select {
		case <-time.After(cns.RetryInterval):
		}
	}
	return nil
}
