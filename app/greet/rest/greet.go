package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-yun/app/greet/rest/internal/config"
	"go-zero-yun/app/greet/rest/internal/handler"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/pkg/cmdkit"
	phandler "go-zero-yun/public/handler"
)

//var cfgFile = flag.String("f", "greet-api.yaml", "the config file")

func main() {
	//加载配置
	cmd := cmdkit.CmdFlagCfg(&config.CfgFile, "greet-api.yaml")
	cmd.Run = webHandler                   //直接运行
	cmd.AddCommand(phandler.GetAliveCmd()) //添加保活命令
	err := cmd.Execute()
	//关闭Client
	defer cmdkit.CloseClient()
	if err != nil {
		panic(fmt.Sprintf("initialize cmd failed: %s", err.Error()))
	}
}

func webHandler(_ *cobra.Command, _ []string) {
	c := config.LoadCfg()
	cmdkit.InitConf(true)

	server := rest.MustNewServer(c.Server.Rest, rest.WithUnauthorizedCallback(phandler.JwtFailRsp))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	handler.OldRegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Server.Rest.Host, c.Server.Rest.Port)
	server.Start()
}
