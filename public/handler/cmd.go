package handler

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-yun/pkg/cmdkit"
	"go-zero-yun/pkg/cobrakit"
	"go-zero-yun/pkg/logkit"
	"go-zero-yun/pkg/prockit"
	"go-zero-yun/public/config"
	"go-zero-yun/public/logic/client/redislgc"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	CmdAliveUse = "alive"
	CmdJobUse   = "job"
	CmdQueueUse = "queue"
	CmdCronUse  = "cron"
)

// GetWgNum 获取命令行协程数
func GetWgNum(args []string) (*sync.WaitGroup, int) {
	var wg sync.WaitGroup
	var grNum int
	if len(args) > 0 {
		grNum, _ = strconv.Atoi(args[0])
	}
	if grNum <= 0 {
		grNum = 3
	}
	wg.Add(grNum)
	return &wg, grNum
}

// GetAliveCmd 获取alive命令
func GetAliveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   CmdAliveUse,
		Short: "run alive",
		Run:   aliveHandler,
	}
}

// aliveHandler alive服务启动
func aliveHandler(_ *cobra.Command, _ []string) {
	fmt.Print("start keep alive for debug\n")
	c := cron.New()
	//每5秒输出当前时间
	_ = c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("alive time: %s\n", time.Now().Format(time.RFC3339))
	})
	c.Start()
	select {}
}

// GetJobCmd 获取job命令
func GetJobCmd(pRun func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:              CmdJobUse,
		Short:            "run job",
		PersistentPreRun: pRun,
	}
}

// GetQueueCmd 获取queue命令
func GetQueueCmd(pRun func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:              CmdQueueUse,
		Short:            "run queue",
		PersistentPreRun: pRun,
	}
}

// GetCronCmd 获取cron命令
func GetCronCmd(pRun func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:              CmdCronUse,
		Short:            "run cron server",
		PersistentPreRun: pRun,
	}
}

func GetCronCmdWithCfg(cfgFile *string, dv string, pRun func(cmd *cobra.Command, args []string)) *cobra.Command {
	c := GetCronCmd(pRun)
	c.Flags().StringVarP(cfgFile, "cron-cfg", "f", "", "go-zero cron config file")
	//如果为空则取默认路径
	if *cfgFile == "" {
		cmdkit.DefaultCfgPath(cfgFile, dv)
	}
	return c
}

func CronInitConf(cmd *cobra.Command, cronCfgFile string) (*cron.Cron, error) {
	if cmd.Context() == nil {
		cmd.SetContext(context.Background())
	}
	//加载配置
	var cronCfg config.CronCfg
	conf.MustLoad(cronCfgFile, &cronCfg)
	prockit.SetWait(time.Duration(cronCfg.QuitWait) * time.Second)
	cmp := make(map[string][][]string)
	for _, v := range cronCfg.Cron {
		if _, ok := cmp[v[1]]; ok {
			cmp[v[1]] = append(cmp[v[1]], v)
		} else {
			cmp[v[1]] = [][]string{v}
		}
	}
	cr := cron.New()
	rCmd := cmd.Parent()
	for _, c := range rCmd.Commands() {
		switch c.Use {
		case CmdJobUse:
			err := CronAddFromCmd(cmd.Context(), cr, c, cmp, "")
			if err != nil {
				return &cron.Cron{}, err
			}
		}
	}
	return cr, nil
}

// CronAddFromCmd cron从cmd添加
func CronAddFromCmd(ctx context.Context, cr *cron.Cron, cmd *cobra.Command, cmp map[string][][]string, pre string) error {
	if pre == "" {
		pre = cmd.Use
	} else {
		pre = fmt.Sprintf("%s/%s", pre, cmd.Use)
	}
	if len(cmd.Commands()) == 0 {
		if _, ok := cmp[pre]; !ok {
			return nil
		}
		for _, vc := range cmp[pre] {
			err := CronAddFunc(ctx, cr, cmd, vc)
			if err != nil {
				return err
			}
		}
		return nil
	}
	for _, c := range cmd.Commands() {
		err := CronAddFromCmd(ctx, cr, c, cmp, pre)
		if err != nil {
			return err
		}
	}
	return nil
}

// CronAddFunc 添加定时任务
func CronAddFunc(ctx context.Context, cr *cron.Cron, cmd *cobra.Command, vc []string) error {
	err := cr.AddFunc(vc[0], func() {
		select {
		case <-prockit.Done():
			return
		default:
		}
		lockKey := fmt.Sprintf("cron:%s", vc[1])
		if len(vc[2:]) > 0 {
			argKey := strings.Join(vc[2:], "|")
			lockKey = fmt.Sprintf("%s:%s", lockKey, argKey)
		}
		//redis分布式锁
		lock := redis.NewRedisLock(redislgc.Data(), lockKey)
		lock.SetExpire(120)
		// 尝试获取锁
		acquire, err := lock.Acquire()
		switch {
		case err != nil:
			logkit.WithType(logkit.LogRedis).Errorf(ctx, "redis lock error:%s", err.Error())
		case acquire:
			defer lock.Release()
		case !acquire:
			// 没有拿到锁直接退出
			logkit.WithType(logkit.LogRun).Infof(ctx, "do not get lock")
			return
		}
		err = cobrakit.RunCmd(cmd, vc[2:])
		if err != nil {
			logkit.WithType(logkit.LogRun).Errorf(ctx, "run cmd error:%s", err.Error())
		}
	})
	return err
}
