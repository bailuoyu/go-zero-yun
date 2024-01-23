package kafkakit

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go-zero-yun/pkg/logkit"
	"time"
	"xorm.io/xorm"
)

type XormFailModel struct {
	Engine *xorm.Engine
}

func (m XormFailModel) Handle(ctx context.Context, msg kafka.Message, retry int, hErr error) {
	// 截断key
	keyRune := []rune(string(msg.Key))
	if len(keyRune) > 200 {
		keyRune = keyRune[:200]
	}
	// 记录头部信息
	hmp := make(map[string]interface{})
	for _, header := range msg.Headers {
		hmp[header.Key] = string(header.Value)
	}
	headers, _ := json.Marshal(hmp)
	log := XormKafkaFailLog{
		Trace:     ctx.Value(logkit.TraceName).(string),
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Offset:    msg.Offset,
		Key:       string(keyRune),
		Value:     string(msg.Value),
		Headers:   string(headers),
		MsgTime:   msg.Time,
		Retry:     retry,
		Err:       hErr.Error(),
	}
	_, err := m.Engine.Context(ctx).Insert(log)
	if err != nil {
		logkit.WithType(logkit.LogKafkaRead).Errorf(ctx, "fail msg record error:%s", err.Error())
	}
}

// XormKafkaFailLog kafka消费失败记录
type XormKafkaFailLog struct {
	Id        int       `xorm:"'id' pk autoincr" json:"id"`                // INT
	Trace     string    `xorm:"'trace'" json:"trace"`                      // comment:链路id;VARCHAR(100)
	Topic     string    `xorm:"'topic'" json:"topic"`                      // default:'';comment:topic;VARCHAR(100)
	Partition int       `xorm:"'partition'" json:"partition"`              // default:0;comment:partition;SMALLINT
	Offset    int64     `xorm:"'offset'" json:"offset"`                    // default:0;comment:offset;BIGINT
	Key       string    `xorm:"'key'" json:"key"`                          // comment:key值;VARCHAR(100)
	Value     string    `xorm:"'value'" json:"Value"`                      // comment:value值;TEXT
	Headers   string    `xorm:"'headers'" json:"headers"`                  // comment:头部信息;JSON
	MsgTime   time.Time `xorm:"'msg_time'" json:"msg_time"`                // comment:消息时间;TIMESTAMP
	Retry     int       `xorm:"retry" json:"retry"`                        // default:0;comment:重试次数;TINYINT
	Err       string    `xorm:"'err'" json:"err"`                          // comment:错误信息;VARCHAR(255)
	Status    int       `xorm:"'status'" json:"status"`                    // default:0;comment:状态;TINYINT
	CreatedAt time.Time `xorm:"'created_at' created <-" json:"created_at"` // default:CURRENT_TIMESTAMP;comment:创建时间;TIMESTAMP
	UpdatedAt time.Time `xorm:"'updated_at' updated <-" json:"updated_at"` // default:CURRENT_TIMESTAMP;comment:更新时间;TIMESTAMP
}

// TableName 表名
func (t XormKafkaFailLog) TableName() string {
	return "kafka_fail_log"
}
