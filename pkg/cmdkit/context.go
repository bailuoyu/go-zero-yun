package cmdkit

import (
	"context"
	"github.com/spf13/cobra"
	"time"
)

// CmdWithCancel 获取cmd的取消
func CmdWithCancel(cmd *cobra.Command) context.CancelFunc {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx2, cancel := context.WithCancel(ctx)
	cmd.SetContext(ctx2)
	return cancel
}

// CmdWithTimeout 获取cmd的超时
func CmdWithTimeout(cmd *cobra.Command, timeout time.Duration) context.CancelFunc {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx2, cancel := context.WithTimeout(ctx, timeout)
	cmd.SetContext(ctx2)
	return cancel
}
