package logkit

import (
	"context"
	"fmt"
	"go-zero-yun/pkg/funckit"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func SetXormLog(engine *xorm.Engine) {
	//xorm是单独设置的engine
	engine.SetLogger(XormLogger{})
}

func SetXormGroupLog(engineGroup *xorm.EngineGroup) {
	//xorm是单独设置的engine
	engineGroup.SetLogger(XormLogger{})
}

//	var ormLog logkit.XormLogger
//	engine.SetLogger(ormLog)

type XormLogger struct {
	CustomLogger log.Logger
	Ctx          context.Context
}

func (c XormLogger) BeforeSQL(context log.LogContext) {
}

func (c XormLogger) AfterSQL(context log.LogContext) {
	runtime := funckit.DurToMic2(context.ExecuteTime)
	//记录日志
	lgr := WithCallerSkip(2)
	if context.Err != nil {
		lgr.WithType(LogXorm).WithRuntime(runtime).
			Errorf(context.Ctx, "error: %s | sql: %s | args: %v", context.Err.Error(), context.SQL, context.Args)
	} else {
		lgr.WithType(LogXorm).WithRuntime(runtime).Infof(context.Ctx, "%s | args: %v", context.SQL, context.Args)
	}

}

func (c XormLogger) Debugf(format string, v ...interface{}) {}

func (c XormLogger) Errorf(format string, v ...interface{}) {}

func (c XormLogger) Infof(format string, v ...interface{}) {}

func (c XormLogger) Warnf(format string, v ...interface{}) {}

func (c XormLogger) Level() log.LogLevel {
	return log.LOG_INFO
}

func (c XormLogger) SetLevel(l log.LogLevel) {}

func (c XormLogger) ShowSQL(show ...bool) {
	fmt.Println(show)
}

func (c XormLogger) IsShowSQL() bool {
	return true
}
