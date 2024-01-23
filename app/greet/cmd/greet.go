package main

import (
	"fmt"
	"go-zero-yun/app/greet/cmd/internal/config"
	"go-zero-yun/app/greet/cmd/internal/svc"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/public/handler"
)

func main() {
	//加载配置
	cmd := cmdkit.CmdFlagCfg(&config.CfgFile, "greet-cmd.yaml")
	cmd.AddCommand(handler.GetAliveCmd()) //添加保活命令

	svc.RegisterJobCmd(cmd)                     //注册job命令
	svc.RegisterQueueCmd(cmd)                   //注册queue命令
	svc.RegisterCronCmd(cmd, "greet-cron.yaml") //注册cron命令

	err := cmd.Execute()
	//关闭Client
	defer cmdkit.CloseClient()
	if err != nil {
		panic(fmt.Sprintf("initialize cmd failed: %s", err.Error()))
	}

}
