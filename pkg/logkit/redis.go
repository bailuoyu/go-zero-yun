package logkit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func SetRedisLog() {
	//redis是全局设置的
	redis.SetLogger(&Redislogger{})
}

type Redislogger struct {
}

func (l *Redislogger) Printf(ctx context.Context, format string, v ...interface{}) {
	//记录日志
	WithCallerSkip(2).WithType(LogRedis).Infof(ctx, format, v...)
}
