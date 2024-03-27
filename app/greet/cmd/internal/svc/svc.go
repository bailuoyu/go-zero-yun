package svc

import (
	"github.com/spf13/cobra"
	"go-zero-yun/app/greet/cmd/internal/config"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/tracekit"
)

// PreFun 前置函数
func PreFun(cmd *cobra.Command, args []string) {
	config.LoadCfgOnce()
	cmdkit.InitConfOnce(false)
	tracekit.CmdStart(cmd, args)
}
