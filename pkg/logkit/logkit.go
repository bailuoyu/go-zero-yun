package logkit

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/funckit"
	"time"
)

type (
	LogField = logx.LogField
)

func WithCallerSkip(skip int) *logger {
	return &logger{
		skip: skip,
	}
}

func WithType(ty string) *logger {
	return &logger{
		ty: ty,
	}
}

func WithDuration(d time.Duration) *logger {
	return &logger{
		runtime: funckit.DurToMic2(d),
	}
}

func WithRuntime(runtime float64) *logger {
	return &logger{
		runtime: runtime,
	}
}

// Debug writes v into access log.
func Debug(ctx context.Context, v ...interface{}) {
	GetLogger(ctx).Debug(v...)
}

// Debugf writes v with format into access log.
func Debugf(ctx context.Context, format string, v ...interface{}) {
	GetLogger(ctx).Debugf(format, v...)
}

// Debugv writes v into access log with json content.
func Debugv(ctx context.Context, v interface{}) {
	GetLogger(ctx).Debugv(v)
}

// Debugw writes msg along with fields into access log.
func Debugw(ctx context.Context, msg string, fields ...LogField) {
	GetLogger(ctx).Debugw(msg, fields...)
}

// Error writes v into error log.
func Error(ctx context.Context, v ...interface{}) {
	GetLogger(ctx).Error(v...)
}

// Errorf writes v with format into error log.
func Errorf(ctx context.Context, format string, v ...interface{}) {
	GetLogger(ctx).Errorf(fmt.Errorf(format, v...).Error())
}

// Errorv writes v into error log with json content.
// No call stack attached, because not elegant to pack the messages.
func Errorv(ctx context.Context, v interface{}) {
	GetLogger(ctx).Errorv(v)
}

// Errorw writes msg along with fields into error log.
func Errorw(ctx context.Context, msg string, fields ...LogField) {
	GetLogger(ctx).Errorw(msg, fields...)
}

// Info writes v into access log.
func Info(ctx context.Context, v ...interface{}) {
	GetLogger(ctx).Info(v...)
}

// Infof writes v with format into access log.
func Infof(ctx context.Context, format string, v ...interface{}) {
	GetLogger(ctx).Infof(format, v...)
}

// Infov writes v into access log with json content.
func Infov(ctx context.Context, v interface{}) {
	GetLogger(ctx).Infov(v)
}

// Infow writes msg along with fields into access log.
func Infow(ctx context.Context, msg string, fields ...LogField) {
	GetLogger(ctx).Infow(msg, fields...)
}

// Slow writes v into slow log.
func Slow(ctx context.Context, v ...interface{}) {
	GetLogger(ctx).Slow(v...)
}

// Slowf writes v with format into slow log.
func Slowf(ctx context.Context, format string, v ...interface{}) {
	GetLogger(ctx).Slowf(format, v...)
}

// Slowv writes v into slow log with json content.
func Slowv(ctx context.Context, v interface{}) {
	GetLogger(ctx).Slowv(v)
}

// Sloww writes msg along with fields into slow log.
func Sloww(ctx context.Context, msg string, fields ...LogField) {
	GetLogger(ctx).Sloww(msg, fields...)
}

// GetLogger returns the logx.Logger with the given ctx and correct caller.
func GetLogger(ctx context.Context) logx.Logger {
	lgr := &logger{
		skip: 1,
		ty:   LogDefault,
	}
	return lgr.GetLogger(ctx)
}
