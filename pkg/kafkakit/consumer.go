package kafkakit

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
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
	Reader        *kafka.Reader
	Handle        func(ctx context.Context, msg kafka.Message) error // bool为标明是否需要重试
	RetryMax      int                                                // 最大重试次数
	RetryInterval time.Duration                                      // 重试间隔
	FailModel     FailModel                                          //消费失败处理函数
	Ctx           context.Context
}

// Close 关闭消费者
func (cns *Consumer) Close() {
	// 计时
	timer := utils.NewElapsedTimer()
	err := cns.Reader.Close()
	if err != nil {
		logkit.WithType(logkit.LogKafkaRead).WithDuration(timer.Duration()).Errorf(cns.Ctx, "reader close err:%s", err.Error())
		fmt.Printf("reader close err: %s,time:%g ms \n", err.Error(), funckit.DurToMic2(timer.Duration()))
	} else {
		logkit.WithType(logkit.LogKafkaRead).WithDuration(timer.Duration()).Info(cns.Ctx, "reader close success")
		fmt.Printf("reader close success,time:%g ms \n", funckit.DurToMic2(timer.Duration()))
	}
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
	// 新建一个带取消的ctx
	ctx, cancel := context.WithCancel(cns.Ctx)
	for {
		select {
		case <-prockit.Done():
			cancel()
		default:
		}
		// kafka消费者自带优雅退出
		m, err := cns.Reader.FetchMessage(ctx)
		if err != nil {
			logkit.WithType(logkit.LogKafkaRead).Errorf(cns.Ctx, "read fetch error:%v", err)
			return err
		}
		msg := Msg{
			Time:      m.Time,
			Topic:     m.Topic,
			Partition: m.Partition,
			Offset:    m.Offset,
			Key:       string(m.Key),
			Value:     string(m.Value),
		}
		// 生成ctx
		htx := cns.handleCtx(trace, m.Partition, m.Offset)
		// 计时
		timer := utils.NewElapsedTimer()
		// 执行handle函数
		jobErr := cns.runHandle(htx, m)
		if jobErr != nil {
			// 记录消费错误日志
			logkit.WithType(logkit.LogKafkaRead).WithDuration(timer.Duration()).
				Errorf(htx, "read partition %d,offset %d error:%s; msg:%+v", m.Partition, m.Offset, jobErr.Error(), msg)
			// 如果是致命错误，则退出
			if jobErr.Code == handler.JobErrFatal {
				return jobErr
			}
		} else {
			// 记录消费日志
			logkit.WithType(logkit.LogKafkaRead).WithDuration(timer.Duration()).
				Infof(htx, "read partition %d,offset %d success; msg:%+v", m.Partition, m.Offset, msg)
		}
		// 确认消费
		err = cns.Reader.CommitMessages(cns.Ctx, m)
		if err != nil {
			logkit.WithType(logkit.LogKafkaRead).Errorf(htx, "read commit partition %d,offset %d error:%s", m.Partition, m.Offset, err.Error())
			return err
		}
	}
}

// handleCtx 设置handle的ctx
func (cns *Consumer) handleCtx(traceStr string, partition int, offset int64) context.Context {
	trace := fmt.Sprintf("%s-%d-%d", traceStr, partition, offset)
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
func (cns *Consumer) runHandle(ctx context.Context, msg kafka.Message) *zerrs.CodeMsg {

	retryNum := 0
	for {
		err := cns.Handle(ctx, msg)
		if err == nil {
			return nil
		}
		jobErr := handler.WrapJobErr(err)
		if jobErr.Code != handler.JobErrRetry || retryNum >= cns.RetryMax {
			//如果没有return说明最终失败，此处应记录到数据库
			if cns.FailModel != nil {
				cns.FailModel.Handle(ctx, msg, retryNum, err)
			}
			return jobErr
		}
		//重试记录info
		logkit.WithType(logkit.LogKafkaRead).Infof(ctx, "read kafka handle error:%s,need retry:%d", err.Error(), retryNum)
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
