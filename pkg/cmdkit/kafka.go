package cmdkit

import (
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/logkit"
	pconf "go-zero-yun/public/config"
	"time"
)

// initKafka 初始化kafka
func initKafka(check bool) {
	pconf.ClientCfg.Kafka = make(map[string]pconf.KafkaCfg)
	for _, v := range pconf.Cfg.Client.Kafka {
		pconf.ClientCfg.Kafka[v.Name] = v
		if check && !IsEnvLocal() { //检查连接
			KafkaConnect(v)
		}
	}
}

// KafkaConnect 检查连接是否可用
func KafkaConnect(v pconf.KafkaCfg) {
	var dialer kafka.Dialer
	mechanism, err := getSASLMechanism(v)
	if err == nil {
		dialer.SASLMechanism = mechanism
	}
	conn, err := dialer.Dial("tcp", v.Brokers[0])
	defer conn.Close()
	if err != nil {
		panic(errors.New(fmt.Sprintf("kafka fatal %s, name: %s", err.Error(), v.Name)))
	}
}

// KafkaWriter kafka生产者
func KafkaWriter(v pconf.KafkaCfg, topic string) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(v.Brokers...),
		Topic:        topic,
		RequiredAcks: kafka.RequireOne, //平衡为1,高性能为0,高安全为-1
		BatchTimeout: 10 * time.Millisecond,
	}
	mechanism, err := getSASLMechanism(v)
	if err != nil {
		return w
	}
	if mechanism == nil {
		return w
	}
	w.Transport = &kafka.Transport{
		SASL: mechanism,
	}
	return w
}

// KafkaReader 获取kafka消费者
func KafkaReader(v pconf.KafkaCfg, topics ...string) *kafka.Reader {
	rCfg := kafkaReaderCfg(v, topics...)
	r := kafka.NewReader(rCfg)
	return r
}

func kafkaReaderCfg(v pconf.KafkaCfg, topics ...string) kafka.ReaderConfig {
	rCfg := kafka.ReaderConfig{
		Brokers:     v.Brokers,
		GroupID:     v.GroupId,
		GroupTopics: topics,
	}
	if rCfg.GroupID == "" { //如果为空则取app名
		rCfg.GroupID = pconf.Cfg.Server.App
	}
	mechanism, err := getSASLMechanism(v)
	if err != nil {
		return rCfg
	}
	if mechanism == nil {
		return rCfg
	}
	rCfg.Dialer = &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	return rCfg
}

// getSASLMechanism 获取sasl验证
func getSASLMechanism(v pconf.KafkaCfg) (sasl.Mechanism, error) {
	var mechanism sasl.Mechanism
	var err error
	switch v.Sasl.Name {
	case "PLAINTEXT":
		return nil, nil
	case "SASL_PLAINTEXT":
		mechanism = plain.Mechanism{
			Username: v.Sasl.Username,
			Password: v.Sasl.Password,
		}
	case "SASL_SCRAM_SHA_256":
		mechanism, err = scram.Mechanism(scram.SHA256, v.Sasl.Username, v.Sasl.Password)
	case "SASL_SCRAM_SHA_512":
		mechanism, err = scram.Mechanism(scram.SHA512, v.Sasl.Username, v.Sasl.Password)
	default:
		return nil, errors.New(fmt.Sprintf("unknown sasl type: %s", v.Sasl.Name))
	}
	if err != nil {
		//记录日志
		logx.Errorw(fmt.Sprintf("kafka sasl err:%s ; name:%s", err.Error(), v.Sasl.Name),
			logx.LogField{Key: logkit.TypeName, Value: logkit.LogKafka})
	}
	return mechanism, err
}
