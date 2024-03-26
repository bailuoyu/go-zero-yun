package queue

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/utils"
	zerrs "github.com/zeromicro/x/errors"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/public/handler"
	"go-zero-yun/public/logic/client/kafkalgc"
	"go-zero-yun/public/model/kafka/core"
	"time"
)

// DemoProduceHandler 生产者
func DemoProduceHandler(cmd *cobra.Command, args []string) {
	//cancel := cmdkit.CmdWithCancel(cmd)
	//go signalkit.CmdElegantExit(cmd, cancel, 5*time.Second)
	// 生产者
	writer := kafkalgc.CoreWriter(core.Log{})
	producer := kafkalgc.GetProducer(cmd.Context(), writer)
	//producer.Writer.Async = true
	//producer.Writer.BatchTimeout = 10 * time.Millisecond
	//defer producer.Close()
	for {
		select {
		case <-cmd.Context().Done():
			fmt.Printf("ctx done time: %s\n", time.Now().Format(time.RFC3339))
			return
		default:

		}
		str := time.Now().Format(time.RFC3339)
		timer := utils.NewElapsedTimer()
		err := producer.Push(kafka.Message{
			Value: []byte(str),
		})
		if err != nil {
			handler.JobError(cmd, err)
			return
		}
		fmt.Printf("produce runtime %g \n", funckit.DurToMic2(timer.Duration()))
		time.Sleep(time.Second * 5)
	}
}

// DemoConsumeHandler 消费者
func DemoConsumeHandler(cmd *cobra.Command, args []string) {
	//cancel := cmdkit.CmdWithCancel(cmd)
	//go signalkit.CmdElegantExit(cmd, cancel, 5*time.Second)
	// 消费者
	reader := kafkalgc.CoreReader(core.Log{})
	consumer := kafkalgc.GetConsumer(cmd.Context(), reader, demoHandle)
	err := consumer.Run()
	fmt.Printf("err: %v", err.Error())
}

func demoHandle(ctx context.Context, msg kafka.Message) error {
	//fmt.Printf("consume:%s; partition %d , offset %d \n", msg.Value, msg.Partition, msg.Offset)
	str := string(msg.Value)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	now := time.Now()
	// 如果延时3秒，则输出
	if now.Second()-t.Second() > 3 {
		fmt.Printf("now:%s - consume:%s; partition %d , offset %d \n", now.Format(time.RFC3339), msg.Value, msg.Partition, msg.Offset)
	}
	return nil
	// 抛出致命错误则直接退出
	return zerrs.New(handler.JobErrFatal, "need return")
	// 抛出重试错误则会重试
	return zerrs.New(handler.JobErrRetry, "need retry")
}
