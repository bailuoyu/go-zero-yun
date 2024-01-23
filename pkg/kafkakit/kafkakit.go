package kafkakit

import "github.com/segmentio/kafka-go"

const MessageIdKey = "message_id"

// GetMessageId 获取MessageId
func GetMessageId(msg kafka.Message) string {
	for _, v := range msg.Headers {
		if v.Key == MessageIdKey {
			return string(v.Value)
		}
	}
	return ""
}
