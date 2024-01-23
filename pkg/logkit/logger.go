package logkit

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/funckit"
	"time"
)

type logger struct {
	skip    int
	runtime float64
	ty      string
}

func (lgr *logger) WithCallerSkip(skip int) *logger {
	lgr.skip = skip
	return lgr
}

func (lgr *logger) WithType(ty string) *logger {
	lgr.ty = ty
	return lgr
}

func (lgr *logger) WithDuration(d time.Duration) *logger {
	lgr.runtime = funckit.DurToMic2(d)
	return lgr
}

func (lgr *logger) WithRuntime(runtime float64) *logger {
	lgr.runtime = runtime
	return lgr
}

// GetLogger returns the logx.Logger with the given ctx and correct caller.
func (lgr *logger) GetLogger(ctx context.Context) logx.Logger {
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
func (lgr *logger) Debug(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Debug(v...)
}

// Debugf writes v with format into access log.
func (lgr *logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Debugf(format, v...)
}

// Debugv writes v into access log with json content.
func (lgr *logger) Debugv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Debugv(v)
}

// Debugw writes msg along with fields into access log.
func (lgr *logger) Debugw(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Debugw(msg, fields...)
}

// Error writes v into error log.
func (lgr *logger) Error(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Error(v...)
}

// Errorf writes v with format into error log.
func (lgr *logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Errorf(fmt.Errorf(format, v...).Error())
}

// Errorv writes v into error log with json content.
// No call stack attached, because not elegant to pack the messages.
func (lgr *logger) Errorv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Errorv(v)
}

// Errorw writes msg along with fields into error log.
func (lgr *logger) Errorw(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Errorw(msg, fields...)
}

// Info writes v into access log.
func (lgr *logger) Info(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Info(v...)
}

// Infof writes v with format into access log.
func (lgr *logger) Infof(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Infof(format, v...)
}

// Infov writes v into access log with json content.
func (lgr *logger) Infov(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Infov(v)
}

// Infow writes msg along with fields into access log.
func (lgr *logger) Infow(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Infow(msg, fields...)
}

// Slow writes v into slow log.
func (lgr *logger) Slow(ctx context.Context, v ...interface{}) {
	lgr.GetLogger(ctx).Slow(v...)
}

// Slowf writes v with format into slow log.
func (lgr *logger) Slowf(ctx context.Context, format string, v ...interface{}) {
	lgr.GetLogger(ctx).Slowf(format, v...)
}

// Slowv writes v into slow log with json content.
func (lgr *logger) Slowv(ctx context.Context, v interface{}) {
	lgr.GetLogger(ctx).Slowv(v)
}

// Sloww writes msg along with fields into slow log.
func (lgr *logger) Sloww(ctx context.Context, msg string, fields ...LogField) {
	lgr.GetLogger(ctx).Sloww(msg, fields...)
}
