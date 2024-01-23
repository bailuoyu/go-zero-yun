package tracekit

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"time"
)

// CliStart queue开始时执行
func CliStart(cmd *cobra.Command, args []string) {
	//设置trace
	CliSetCtx(cmd)
	//记录开始日志
	logkit.WithType(logkit.LogRunStart).Infof(cmd.Context(), "args: %v", args)
}

// CliSetCtx 设置job的ctx
func CliSetCtx(cmd *cobra.Command) context.Context {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, logkit.StartTimeName, time.Now().UnixMicro())
	trace := funckit.RandomStr(24, 1)
	ctx = context.WithValue(ctx, logkit.TraceName, trace)
	ctx2 := logx.WithFields(ctx,
		logx.LogField{
			Key:   logkit.TraceName,
			Value: trace,
		},
		logx.LogField{
			Key:   logkit.RouteName,
			Value: GetCmdRoute(cmd),
		},
	)
	cmd.SetContext(ctx2)
	return ctx2
}

// GetCmdRoute 获取cmd的route
func GetCmdRoute(cmd *cobra.Command) string {
	route := cmd.Use
	for p := cmd.Parent(); p != nil; p = p.Parent() {
		if p.Use == "" {
			break
		}
		route = fmt.Sprintf("%s/%s", p.Use, route)
	}
	return route
}
