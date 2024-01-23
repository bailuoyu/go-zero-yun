// Package plugin 通用插件工厂体系，提供插件注册和装载，主要用于需要通过动态配置加载生成具体插件的情景，不需要配置生成的插件，可直接注册到具体的插件包里面(如codec) 不需要注册到这里。
package plugin

import (
	"errors"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

const DefaultName = "Default"

var (
	// SetupTimeout 每个插件初始化最长超时时间，如果某个插件确实需要加载很久，可以自己修改这里的值
	SetupTimeout = 3 * time.Second

	// MaxPluginSize  最大插件个数
	MaxPluginSize = 1000
)

var (
	pluginTypes = make(map[string]Factory)            // plugin type => { plugin name => plugin factory }
	plugins     = make(map[string]map[string]Factory) // plugin type => { plugin name => plugin factory }
	done        = make(chan struct{})                 // 插件初始化完成通知channel
)

// WaitForDone 挂住等待所有插件初始化完成，可自己设置超时时间。
func WaitForDone(timeout time.Duration) bool {
	select {
	case <-done:
		return true
	case <-time.After(timeout):
	}
	return false
}

// Factory 插件工厂统一抽象，外部插件需要实现该接口，通过该工厂接口生成具体的插件并注册到具体的插件类型里面。
type Factory interface {
	// Type 插件的类型 如 selector log config tracing
	Type() string
	// Setup 根据配置项节点装载插件，需要用户自己先定义好具体插件的配置数据结构
	Setup(name string, dec Decoder) error
}

// Decoder 节点配置解析器。
type Decoder interface {
	Decode(cfg interface{}) error // 输入参数为自定义的配置数据结构
}

// MapstructureDecoder yaml节点配置解析器。
type MapstructureDecoder struct {
	Node interface{}
}

// Decode 解析mapstructure node配置。
func (d MapstructureDecoder) Decode(cfg interface{}) error {
	if d.Node == nil {
		return errors.New("mapstructure node empty")
	}
	return mapstructure.Decode(d.Node, cfg)
}

// RegisterType 注册插件工厂类型
func RegisterType(f Factory) {
	if _, ok := pluginTypes[f.Type()]; !ok {
		pluginTypes[f.Type()] = f
	}
}

// GetType 根据插件类型，插件名字获取插件工厂。
func GetType(typ string) Factory {
	if _, ok := pluginTypes[typ]; ok {
		return pluginTypes[typ]
	} else {
		return nil
	}
}

// Register 注册插件工厂 可自己指定插件名，支持相同的实现，不同的配置注册不同的工厂实例。
// 弃用
func Register(name string, f Factory) {
	factories, ok := plugins[f.Type()]
	if !ok {
		plugins[f.Type()] = map[string]Factory{
			name: f,
		}
		return
	}
	factories[name] = f
}

// Get 根据插件类型，插件名字获取插件工厂。
// 弃用
func Get(typ string, name string) Factory {
	factories, ok := plugins[typ]
	if !ok {
		return nil
	}
	return factories[name]
}

// Config 插件统一配置 plugin type => { plugin name => plugin config } 。
// type Config map[string]map[string]interface mapstructure
type Config map[string]map[string]interface{}

// Setup 通过配置生成并装载具体插件。
func (c Config) Setup() error {
	var (
		pluginChan  = make(chan Info, MaxPluginSize) // 初始化插件队列，方便后面按顺序逐个加载插件
		setupStatus = make(map[string]bool)          // 插件初始化状态，plugin key => true初始化完成 false未初始化
	)

	// 从框架配置文件中逐个取出插件并放进channel队列中
	for typ, facts := range c {
		fact := GetType(typ)
		if fact == nil {
			return fmt.Errorf("plugin type %s no registered or imported, do not configure", typ)
		}
		for name, cfg := range facts {
			newFact := fact
			p := Info{
				factory: newFact,
				typ:     typ,
				name:    name,
				cfg:     cfg,
			}
			select {
			case pluginChan <- p:
			default:
				return fmt.Errorf("plugin number exceed max limit:%d", len(pluginChan))
			}
			setupStatus[p.Key()] = false
		}
	}

	// 从channel队列中取出插件并初始化
	num := len(pluginChan)
	for num > 0 {
		for i := 0; i < num; i++ {
			p := <-pluginChan
			// 先判断当前插件依赖的其他插件是否已经初始化完成
			if deps, err := p.Depends(setupStatus); err != nil {
				return err
			} else if deps { // 被依赖的插件还未初始化，将当前插件移到channel末尾
				pluginChan <- p
				continue
			}
			if err := p.Setup(); err != nil {
				return err
			}
			setupStatus[p.Key()] = true
		}
		if len(pluginChan) == num { // 循环依赖导致无插件可以初始化，返回失败
			return fmt.Errorf("cycle depends, not plugin is setup")
		}
		num = len(pluginChan)
	}

	// 发出插件初始化完成通知，个别业务逻辑需要依赖插件完成才能继续往下执行
	select {
	case <-done: // 已经close过了
	default:
		close(done)
	}
	return nil
}

// Depender 依赖接口，由具体实现插件决定是否有依赖其他插件, 需要保证被依赖的插件先初始化完成。
type Depender interface {
	// DependsOn 假如一个插件依赖另一个插件，则返回被依赖的插件的列表：数组元素为 type-name 如 [ "selector-polaris" ]
	DependsOn() []string
}

// Info 插件信息。
type Info struct {
	factory Factory
	typ     string
	name    string
	cfg     interface{}
}

// Depends 判断是否有依赖的插件未初始化过。
// 输入参数为所有插件的初始化状态。
// 输出参数bool true被依赖的插件未初始化完成，仍有依赖，false没有依赖其他插件或者被依赖的插件已经初始化完成
func (p *Info) Depends(setupStatus map[string]bool) (bool, error) {
	deps, ok := p.factory.(Depender)
	if !ok { // 该插件不依赖任何其他插件
		return false, nil
	}
	depends := deps.DependsOn()
	for _, name := range depends {
		if name == p.Key() {
			return false, fmt.Errorf("cannot depends on self")
		}
		setup, ok := setupStatus[name]
		if !ok {
			return false, fmt.Errorf("depends plugin %s not exists", name)
		}
		if !setup {
			return true, nil
		}
	}
	return false, nil
}

// Setup 初始化单个插件。
func (p *Info) Setup() error {
	var (
		ch  = make(chan struct{})
		err error
	)
	go func() {
		//err = p.factory.Setup(p.name, &MapstructureDecoder{Node: &p.cfg})
		err = p.factory.Setup(p.name, MapstructureDecoder{Node: p.cfg})
		close(ch)
	}()
	select {
	case <-ch:
	case <-time.After(SetupTimeout):
		return fmt.Errorf("setup plugin %s timeout", p.Key())
	}
	if err != nil {
		return fmt.Errorf("setup plugin %s error: %v", p.Key(), err)
	}
	return nil
}

// Key 插件的唯一索引：type-name 。
func (p *Info) Key() string {
	return fmt.Sprintf("%s-%s", p.typ, p.name)
}
