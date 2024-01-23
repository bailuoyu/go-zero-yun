package svc

import (
	"github.com/spf13/cobra"
	"go-zero-yun/app/greet/cmd/internal/handler/job"
	"go-zero-yun/public/handler"
)

// RegisterJobCmd 注册job命令
func RegisterJobCmd(cmd *cobra.Command) {
	c := handler.GetJobCmd(PreFun)
	c.AddCommand(JobCommands()...)
	cmd.AddCommand(c)
}

// JobCommands 注册job
func JobCommands() []*cobra.Command {
	return []*cobra.Command{
		{Use: "test", Run: job.TestHandler},     // 测试用
		{Use: "signal", Run: job.SignalHandler}, // 测试信号用
	}
}
