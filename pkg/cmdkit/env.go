package cmdkit

import pconf "go-zero-yun/public/config"

// IsEnvDev 判断是否为开发环境
func IsEnvDev() bool {
	switch pconf.Cfg.Global.EnvName {
	case "local", "dev":
		return true
	default:
		return false
	}
}

// IsEnvProd 判断是否为生产环境
func IsEnvProd() bool {
	switch pconf.Cfg.Global.EnvName {
	case "prod", "pre":
		return true
	default:
		return false
	}
}

// IsEnvLocal 判断是否为本地环境
func IsEnvLocal() bool {
	switch pconf.Cfg.Global.EnvName {
	case "local":
		return true
	default:
		return false
	}
}
