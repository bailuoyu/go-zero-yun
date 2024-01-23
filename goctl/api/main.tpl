package main

import (
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/public/cmd/alive"
	"fmt"
    "github.com/spf13/cobra"
    "github.com/zeromicro/go-zero/core/logx"
    "os"

	{{.importPackages}}
)

//var cfgFile = flag.String("f", "{{.serviceName}}.yaml", "the config file")

func main() {
    cmd := cmdkit.CmdFlagParse(&config.CfgFile, "{{.serviceName}}.yaml")
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

	server := rest.MustNewServer(c.Server.Rest)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Server.Rest.Host, c.Server.Rest.Port)
	server.Start()
}
