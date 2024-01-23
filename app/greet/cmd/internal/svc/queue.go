package svc

import (
	"github.com/spf13/cobra"
	"go-zero-yun/app/greet/cmd/internal/handler/queue"
	"go-zero-yun/public/handler"
)

// RegisterQueueCmd 注册queue命令
func RegisterQueueCmd(cmd *cobra.Command) {
	c := handler.GetQueueCmd(PreFun)
	c.AddCommand(QueueCommands()...)
	cmd.AddCommand(c)
}

// QueueCommands 注册queue
func QueueCommands() []*cobra.Command {
	return []*cobra.Command{
		{Use: "demo_produce", Run: queue.DemoProduceHandler},
		{Use: "demo_consume", Run: queue.DemoConsumeHandler},
		{Use: "demo2_produce", Run: queue.Demo2ProduceHandler},
		{Use: "demo2_consume", Run: queue.Demo2ConsumeHandler},
	}
}
