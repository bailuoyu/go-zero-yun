package rabbitmqkit

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/utils"
	zerrs "github.com/zeromicro/x/errors"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"go-zero-yun/pkg/prockit"
	"go-zero-yun/public/handler"
	"time"
)

const (
	DefaultRetryMax      = 3
	DefaultRetryInterval = 100 * time.Millisecond
)

// Consumer 消费者
type Consumer struct {
	Client
	Handle        func(ctx context.Context, msg amqp.Delivery) error // bool为标明是否需要重试
	Queue         string
	NoWait        bool
	RetryMax      int           // 最大重试次数
	RetryInterval time.Duration // 重试间隔
	FailModel     FailModel
	Ctx           context.Context
}

// Close 关闭生产者
func (cns *Consumer) Close() {
	// 计时
	timer := utils.NewElapsedTimer()
	err1 := cns.CloseChannel()
	runtime1 := funckit.DurToMic2(timer.Duration())
	if err1 != nil {
		logkit.WithType(logkit.LogKafkaRead).WithRuntime(runtime1).Errorf(cns.Ctx, "reader close err: %s", err1.Error())
		fmt.Printf("reader close err: %s,time:%g ms \n", err1.Error(), runtime1)
	}
	err2 := cns.CloseCon()
	runtime2 := funckit.DurToMic2(timer.Duration())
	if err2 != nil {
		logkit.WithType(logkit.LogKafkaRead).WithRuntime(runtime2).Errorf(cns.Ctx, "reader close err: %s", err2.Error())
		fmt.Printf("reader close err: %s,time:%g ms \n", err2.Error(), runtime2)
	}
	if err1 == nil && err2 == nil {
		logkit.WithType(logkit.LogKafkaRead).WithRuntime(runtime2).Infof(cns.Ctx, "reader close success")
		fmt.Printf("reader close success,time:%g ms \n", runtime2)
	}
}

// CloseCon 关闭连接
func (cns *Consumer) CloseCon() error {
	return cns.Con.Close()
}

// CloseChannel 关闭通道
func (cns *Consumer) CloseChannel() error {
	return cns.Channel.Close()
}

// Run 运行消费者
func (cns *Consumer) Run() error {
	//检查并设置配置默认参数
	if cns.RetryMax == 0 {
		cns.RetryMax = DefaultRetryMax
	}
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
	consumerTag := funckit.RandomStr(8, 1)
	msgc, err := cns.Channel.Consume(cns.Queue, consumerTag, false, false, false, cns.NoWait, nil)
	if err != nil {
		logkit.WithType(logkit.LogRabbitmqRead).Errorf(cns.Ctx, "read fetch error:%v", err)
		return err
	}

	for {
		select {
		case <-prockit.Done(): //判断服务是否结束
			return context.Canceled
		case m := <-msgc:
			msg := Msg{
				Time:       m.Timestamp,
				Exchange:   m.Exchange,
				RoutingKey: m.RoutingKey,
				MessageId:  m.MessageId,
				Body:       string(m.Body),
			}
			// 生成ctx
			htx := cns.handleCtx(trace, m.ConsumerTag, m.DeliveryTag)
			// 计时
			timer := utils.NewElapsedTimer()
			// 执行handle函数
			jobErr := cns.runHandle(htx, m)
			if jobErr != nil {
				// 记录消费错误日志
				logkit.WithType(logkit.LogRabbitmqRead).WithDuration(timer.Duration()).
					Errorf(htx, "read tag %d,MessageId %s error:%s; msg:%+v", m.DeliveryTag, m.MessageId, jobErr.Error(), msg)
				if jobErr.Code == handler.JobErrFatal {
					return jobErr
				}
			} else {
				// 记录消费日志
				logkit.WithType(logkit.LogRabbitmqRead).WithDuration(timer.Duration()).
					Infof(htx, "read tag %d,MessageId %s success; msg:%+v", m.DeliveryTag, m.MessageId, msg)
			}
			// 确认消费
			err = m.Ack(false)
			if err != nil {
				logkit.WithType(logkit.LogRabbitmqRead).Errorf(htx, "read commit tag %d,MessageId %s error:%s", m.DeliveryTag, m.MessageId, err.Error())
				return err
			}
		}
	}
}

// handleCtx 设置handle的ctx
func (cns *Consumer) handleCtx(traceStr string, consumerTag string, deliveryTag uint64) context.Context {
	trace := fmt.Sprintf("%s-%s-%d", traceStr, consumerTag, deliveryTag)
	htx := context.WithValue(cns.Ctx, logkit.TraceName, trace)
	htx = logx.WithFields(htx,
		logx.LogField{
			Key:   logkit.TraceName,
			Value: trace,
		},
		logx.LogField{
			Key:   logkit.SpanName,
			Value: traceStr,
		},
	)
	return htx
}

// runHandle 执行handle
func (cns *Consumer) runHandle(ctx context.Context, msg amqp.Delivery) *zerrs.CodeMsg {
	retryNum := 0
	for {
		err := cns.Handle(ctx, msg)
		if err == nil {
			return nil
		}
		jobErr := handler.WrapJobErr(err)
		if jobErr.Code != handler.JobErrRetry || retryNum >= cns.RetryMax {
			// 此处应记录到数据库
			if cns.FailModel != nil {
				cns.FailModel.Handle(ctx, msg, retryNum, err)
			}
			return jobErr
		}
		//重试则记录info
		logkit.WithType(logkit.LogRabbitmqRead).
			Info(ctx, "read Rabbitmq handle error:%s,need retry:%d", err.Error(), retryNum)
		retryNum++
		// 如果需要重试，就等待一段时间后再次执行
		t := time.NewTimer(cns.RetryInterval)
		select {
		case <-cns.Ctx.Done(): //判断服务是否结束
			return &zerrs.CodeMsg{
				Code: -1,
				Msg:  cns.Ctx.Err().Error(),
			}
		case <-ctx.Done(): //判断协程是否结束
			return &zerrs.CodeMsg{
				Code: -1,
				Msg:  ctx.Err().Error(),
			}
		case <-t.C:
			// retry
		}
	}
}
