package cmdkit

import (
	"fmt"
	"github.com/spf13/cobra"
	pconf "go-zero-yun/public/config"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	confOnce sync.Once
)

// InitConf 初始化配置
func InitConf(check bool) {
	InitClient(check)
	// 初始化Plugin配置
	if pconf.Cfg.Plugin != nil {
		if err := pconf.Cfg.Plugin.Setup(); err != nil {
			panic(fmt.Errorf("Fatal load plugin: %s \n", err))
		}
	}
}

// InitConfOnce 防止协程中多次加载
func InitConfOnce(check bool) {
	// 防止协程中多次加载
	confOnce.Do(func() {
		InitConf(check)
	})
}

// CmdFlagCfg 获取cmd配置
func CmdFlagCfg(cfgFile *string, dv string) *cobra.Command {
	cmd := &cobra.Command{
		Short: "Common Service Platform",
	}
	cmd.PersistentFlags().StringVarP(cfgFile, "config", "c", "", "go-zero config file")
	//如果为空则取默认路径
	if *cfgFile == "" {
		DefaultCfgPath(cfgFile, dv)
	}
	return cmd
}

// DefaultCfgPath 获取默认路径，防止出错
func DefaultCfgPath(cfgFile *string, dv string) {
	pathDir, _ := filepath.Abs(os.Args[0])
	index := strings.LastIndex(pathDir, string(os.PathSeparator))
	pathDir = pathDir[:index]
	*cfgFile = path.Join(pathDir, dv)
}
