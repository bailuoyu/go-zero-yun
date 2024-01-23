package cos

import (
	"fmt"
	"go-zero-yun/plugin"
)

const pluginType = "Cos"

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
	var cfg Config
	err := dec.Decode(&cfg)
	if err != nil {
		return err
	}
	configMp[name] = cfg
	fmt.Printf("plugin %s:%s config: %+v \n", pluginType, name, cfg)
	return err
}
