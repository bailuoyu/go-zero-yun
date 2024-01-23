package queue

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	zerrs "github.com/zeromicro/x/errors"
	"go-zero-yun/pkg/logkit"
	"go-zero-yun/public/handler"
	"go-zero-yun/public/logic/client/rabbitmqlgc"
	quem "go-zero-yun/public/model/rabbitmq/core/queue"
	"time"
)

// Demo2ProduceHandler 测试,第一个参数为协程数
func Demo2ProduceHandler(cmd *cobra.Command, args []string) {
	//cancel := cmdkit.CmdWithCancel(cmd)
	//go signalkit.CmdElegantExit(cmd, cancel, 5*time.Second)
	// 生产者
	client := rabbitmqlgc.Core()
	producer := rabbitmqlgc.GetProducer(cmd.Context(), client, quem.Log{})
	defer producer.Close()
	for {
		select {
		case <-cmd.Context().Done():
			fmt.Printf("ctx done time: %s\n", time.Now().Format(time.RFC3339))
			return
		default:

		}
		str := time.Now().Format(time.RFC3339)
		logkit.Infof(cmd.Context(), "push:%v\n", str)
		//timer := utils.NewElapsedTimer()
		err := producer.Push(amqp.Publishing{
			Body: []byte(str),
		})
		if err != nil {
			//logkit.Errorf(cmd.Context(), "%v", err)
			handler.JobError(cmd, err)
			return
		}
		//fmt.Printf("produce runtime %g", funckit.DurToMic2(timer.Duration()))
		time.Sleep(time.Second * 5)
	}

	// 延时消息生产者
	//client2 := rabbitmqlgc.Core()
	//producer2 := rabbitmqlgc.GetProducer(cmd.Context(), client2, exchange.DelayLog{})
	//defer producer2.Close()
	//go func() {
	//	for {
	//		str := time.Now().Format(time.RFC3339)
	//		logkit.Infof(cmd.Context(), "push:%v\n", str)
	//		msg := amqp.Publishing{
	//			Body: []byte(str),
	//		}
	//		rabbitmqkit.DelayMsg(&msg, 30*time.Second) //延时30秒
	//		err := producer2.Push(msg)
	//		if err != nil {
	//			handler.JobError(cmd, err)
	//			return
	//		}
	//		time.Sleep(time.Second * 5)
	//	}
	//}()
}

func Demo2ConsumeHandler(cmd *cobra.Command, args []string) {
	//cancel := cmdkit.CmdWithCancel(cmd)
	//go signalkit.CmdElegantExit(cmd, cancel, 5*time.Second)
	// 消费者
	client := rabbitmqlgc.Core()
	consumer := rabbitmqlgc.GetConsumer(cmd.Context(), client, quem.Log{}, demo2Handle)
	err := consumer.Run()
	fmt.Printf("err: %v", err.Error())
}

func demo2Handle(ctx context.Context, msg amqp.Delivery) error {
	fmt.Printf("consume:%s\n", msg.Body)
	return nil
	// 抛出致命错误则直接退出
	return zerrs.New(handler.JobErrFatal, "need return")
	// 抛出重试错误则会重试
	return zerrs.New(handler.JobErrRetry, "need retry")
}
