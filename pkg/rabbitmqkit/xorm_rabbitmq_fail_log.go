package rabbitmqkit

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-zero-yun/pkg/logkit"
	"time"
	"xorm.io/xorm"
)

type XormFailModel struct {
	Engine *xorm.Engine
}

func (m XormFailModel) Handle(ctx context.Context, msg amqp.Delivery, retry int, hErr error) {
	headers, _ := json.Marshal(msg.Headers)
	log := XormRabbitmqFailLog{
		Trace:      ctx.Value(logkit.TraceName).(string),
		Exchange:   msg.Exchange,
		RoutingKey: msg.RoutingKey,
		MessageId:  msg.MessageId,
		Body:       string(msg.Body),
		Headers:    string(headers),
		MsgTime:    msg.Timestamp,
		Retry:      retry,
		Err:        hErr.Error(),
	}
	_, err := m.Engine.Context(ctx).Insert(log)
	if err != nil {
		logkit.WithType(logkit.LogRabbitmqRead).Errorf(ctx, "fail msg record error:%s", err.Error())
	}
}

// XormRabbitmqFailLog Rabbitmq消费失败记录
type XormRabbitmqFailLog struct {
	Id         int       `xorm:"'id' pk autoincr" json:"id"`                // INT
	Trace      string    `xorm:"'trace'" json:"trace"`                      // comment:链路id;VARCHAR(100)
	Exchange   string    `xorm:"'exchange'" json:"exchange"`                // default:'';comment:exchange;VARCHAR(100)
	RoutingKey string    `xorm:"'routing_key'" json:"routing_key"`          // default:0;comment:routing_key;VARCHAR(100)
	MessageId  string    `xorm:"'message_id'" json:"message_id"`            // comment:message_id值;VARCHAR(100)
	Body       string    `xorm:"'body'" json:"body"`                        // comment:body值;TEXT
	Headers    string    `xorm:"'headers'" json:"headers"`                  // comment:头部信息;JSON
	MsgTime    time.Time `xorm:"'msg_time'" json:"msg_time"`                // comment:消息时间;TIMESTAMP
	Retry      int       `xorm:"retry" json:"retry"`                        // default:0;comment:重试次数;TINYINT
	Err        string    `xorm:"'err'" json:"err"`                          // comment:错误信息;VARCHAR(255)
	Status     int       `xorm:"'status'" json:"status"`                    // default:0;comment:状态;TINYINT
	CreatedAt  time.Time `xorm:"'created_at' created <-" json:"created_at"` // default:CURRENT_TIMESTAMP;comment:创建时间;TIMESTAMP
	UpdatedAt  time.Time `xorm:"'updated_at' updated <-" json:"updated_at"` // default:CURRENT_TIMESTAMP;comment:更新时间;TIMESTAMP
}

// TableName 表名
func (t XormRabbitmqFailLog) TableName() string {
	return "rabbitmq_fail_log"
}
