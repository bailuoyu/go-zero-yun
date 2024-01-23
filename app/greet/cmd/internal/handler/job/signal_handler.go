package job

import (
	"github.com/spf13/cobra"
	"go-zero-yun/app/greet/cmd/internal/logic/joblgc"
)

// SignalHandler 测试
func SignalHandler(cmd *cobra.Command, args []string) {
	// 5秒的优雅停止
	//cancel := cmdkit.CmdWithCancel(cmd)
	//go signalkit.CmdElegantExit(cmd, cancel, 5*time.Second)
	joblgc.SignalLogic(cmd, args)
}
