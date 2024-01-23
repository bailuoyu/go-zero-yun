package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/x/errors"
	"go-zero-yun/pkg/logkit"
	"strconv"
	"time"
)

const (
	JobErrFatal = -9
	JobErrRetry = -2
)

// JobResult 返回结果
func JobResult(cmd *cobra.Command, err error) {
	if err == nil {
		JobSuccess(cmd)
	} else {
		JobError(cmd, err)
	}
}

// JobSuccess 脚本打印成功信息
func JobSuccess(cmd *cobra.Command, raw ...string) {
	var msg string
	if len(raw) > 0 {
		msg = fmt.Sprintf("success raw:%+v", raw)
	} else {
		msg = "success"
	}
	logkit.WithType(logkit.LogRunEnd).WithRuntime(getJobExcTime(cmd)).Info(cmd.Context(), msg)
	fmt.Printf("%s\n", msg)
}

// JobError 打印定时脚本错误，方便监控
func JobError(cmd *cobra.Command, err error, raw ...string) {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	if len(raw) > 0 {
		msg = fmt.Sprintf("err %s;raw:%v", msg, raw)
	}
	logkit.WithType(logkit.LogRunEnd).WithRuntime(getJobExcTime(cmd)).Error(cmd.Context(), msg)
	fmt.Printf("%s\n", msg)
}

// getExcTime 获取执行时间
func getJobExcTime(cmd *cobra.Command) float64 {
	start := cmd.Context().Value(logkit.StartTimeName)
	startStr := fmt.Sprintf("%d", start)
	startInt64, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil || startInt64 == 0 {
		return -1
	}
	excTime := time.Now().UnixMicro() - startInt64
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(excTime)/1000), 64)
	return value
}

// WrapJobErr 解析错误函数
func WrapJobErr(v any) *errors.CodeMsg {
	switch data := v.(type) {
	case *errors.CodeMsg:
		return data
	case errors.CodeMsg:
		return &data
	//case *status.Status:
	//	return errors.CodeMsg{
	//		Code: int(data.Code()),
	//		Msg:  data.Message(),
	//	}
	case error:
		return &errors.CodeMsg{
			Code: -1,
			Msg:  data.Error(),
		}
	}
	return &errors.CodeMsg{}
}
