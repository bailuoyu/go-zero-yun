package config

// Global 全局变量
type Global struct {
	Namespace string `json:"Namespace,default=Development,options=Development|Testing|Product"`
	EnvName   string `json:"EnvName,default=dev,options=local|dev|test|prod"`
}
