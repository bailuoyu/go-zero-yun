package weibo

import (
	"go-zero-yun/plugin"
)

const pluginType = "WEIBO"

// Factory 构建工厂
type Factory struct{}

func init() {
	plugin.RegisterType(Factory{})
}

// Type 类型
func (f Factory) Type() string {
	return pluginType
}

// Setup 设置
func (f Factory) Setup(name string, dec plugin.Decoder) error {
	return nil
}
