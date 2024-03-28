package logkit

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/funckit"
	"time"
)

type Logger struct {
	skip    int
	runtime float64
	ty      string
}

func (lgr *Logger) WithCallerSkip(skip int) *Logger {
	lgr.skip = skip
	return lgr
}

func (lgr *Logger) WithType(ty string) *Logger {
	lgr.ty = ty
	return lgr
}

func (lgr *Logger) WithDuration(d time.Duration) *Logger {
	lgr.runtime = funckit.DurToMic2(d)
	return lgr
}

func (lgr *Logger) WithRuntime(runtime float64) *Logger {
	lgr.runtime = runtime
	return lgr
}

// GetLogger returns the logx.Logger with the given ctx and correct caller.
func (lgr *Logger) GetLogger(ctx context.Context) logx.Logger {
	xlgr := logx.WithContext(ctx)
	if lgr.skip < 1 {
		lgr.skip = 1
	}
	xlgr = xlgr.WithCallerSkip(lgr.skip)
	if lgr.ty == "" {
		lgr.ty = LogDefault
	}
	xlgr = xlgr.WithFields(LogField{Key: TypeName, Value: lgr.ty})
	if lgr.runtime > 0 {
		xlgr = xlgr.WithFields(LogField{Key: RuntimeName, Value: lgr.runtime})
	}
	return xlgr
}

// Debug writes v into access log.
func (lgr *Logger) Debug(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Debug(v...)
}

// Debugf writes v with format into access log.
func (lgr *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Debugf(format, v...)
}

// Debugv writes v into access log with json content.
func (lgr *Logger) Debugv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Debugv(v)
}

// Debugw writes msg along with fields into access log.
func (lgr *Logger) Debugw(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Debugw(msg, fields...)
}

// Error writes v into error log.
func (lgr *Logger) Error(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Error(v...)
}

// Errorf writes v with format into error log.
func (lgr *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Errorf(fmt.Errorf(format, v...).Error())
}

// Errorv writes v into error log with json content.
// No call stack attached, because not elegant to pack the messages.
func (lgr *Logger) Errorv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Errorv(v)
}

// Errorw writes msg along with fields into error log.
func (lgr *Logger) Errorw(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Errorw(msg, fields...)
}

// Info writes v into access log.
func (lgr *Logger) Info(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Info(v...)
}

// Infof writes v with format into access log.
func (lgr *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Infof(format, v...)
}

// Infov writes v into access log with json content.
func (lgr *Logger) Infov(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Infov(v)
}

// Infow writes msg along with fields into access log.
func (lgr *Logger) Infow(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Infow(msg, fields...)
}

// Slow writes v into slow log.
func (lgr *Logger) Slow(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Slow(v...)
}

// Slowf writes v with format into slow log.
func (lgr *Logger) Slowf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Slowf(format, v...)
}

// Slowv writes v into slow log with json content.
func (lgr *Logger) Slowv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Slowv(v)
}

// Sloww writes msg along with fields into slow log.
func (lgr *Logger) Sloww(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Sloww(msg, fields...)
}
