package monkit

import (
	"context"
	"github.com/zeromicro/go-zero/core/timex"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"strings"
	"time"
)

const mongoAddrSep = ","

// FormatAddr formats mongo hosts to a string.
func FormatAddr(hosts []string) string {
	return strings.Join(hosts, mongoAddrSep)
}

func logDuration(ctx context.Context, name, method string, startTime time.Duration, err error) {
	duration := timex.Since(startTime)
	runtime := funckit.DurToMic2(duration)
	if err != nil {
		logkit.WithType(logkit.LogMongo).WithRuntime(runtime).
			Errorf(ctx, "mongo(%s) - %s - fail(%s)", name, method, err.Error())
	} else {
		logkit.WithType(logkit.LogMongo).WithRuntime(runtime).
			Infof(ctx, logkit.LogMongo, "mongo(%s) - %s - ok", name, method)
	}
}
