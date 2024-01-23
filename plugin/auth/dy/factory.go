package dy

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/plugin"
)

const pluginType = "DY"

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
	v, _ := jsoniter.MarshalToString(cfg)
	logx.Infof("plugin %s:%s config: %s", pluginType, name, v)
	return err
}
