package config

// KafkaCfg Kafka配置结构体
type KafkaCfg struct {
	Name    string   `json:"Name"`
	Brokers []string `json:"Brokers"`
	Sasl    struct {
		Name     string `json:"Name"`
		Username string `json:"Username,optional"`
		Password string `json:"Password,optional"`
	} `json:"Sasl,optional"`
	GroupId string `json:"GroupId,optional"`
}
