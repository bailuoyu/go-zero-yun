package svc

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-zero-yun/public/handler"
)

var cronCfgFile string

// RegisterCronCmd 注册cron命令
func RegisterCronCmd(cmd *cobra.Command, dv string) {
	c := handler.GetCronCmdWithCfg(&cronCfgFile, dv, PreFun)
	c.Run = cronHandler
	cmd.AddCommand(c)
}

// cronHandler 启动定时任务
func cronHandler(cmd *cobra.Command, _ []string) {
	fmt.Printf("%s starting\n", cmd.Short)

	cron, err := handler.CronInitConf(cmd, cronCfgFile)
	if err != nil {
		panic(fmt.Sprintf("initialize cron failed: %s", err.Error()))
	}

	// 测试输出时间
	//_ = cron.AddFunc("*/10 * * * * *", func() {
	//	fmt.Printf("cron time: %s\n", time.Now().Format(time.RFC3339))
	//})

	cron.Run()

}
