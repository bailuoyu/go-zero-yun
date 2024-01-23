package prockit

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/pkg/funckit"
	"go-zero-yun/pkg/logkit"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultWait = 15 * time.Second
)

var done = make(chan struct{})

var wait = defaultWait
var mutex sync.Mutex

func init() {
	go func() {
		// https://golang.org/pkg/os/signal/#Notify
		signals := make(chan os.Signal, 1)
		//signal.Notify(signals, syscall.SIGTERM)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
		ctx := context.Background()
	signalLoop:
		for {
			v := <-signals
			switch v {
			case syscall.SIGTERM, syscall.SIGINT:
				logx.Infof("Got registered signal:%s", v)
				select {
				case <-done:
					// already closed
				default:
					close(done)
				}
				break signalLoop
			default:
				logx.Infof("Got unregistered signal:%s", v)
			}
		}
		runtime := funckit.DurToMic2(wait)
		logkit.Infof(ctx, "wait up to %g ms before exiting", runtime)
		fmt.Printf("wait up to %g ms before exiting \n", runtime)
		<-time.After(wait)
		// 超时退出则记录错误
		logkit.WithRuntime(runtime).Errorf(ctx, "signal timeout, exit \n")
		fmt.Errorf("signal timeout %g ms, exit", runtime)
		os.Exit(0)
	}()
}

// Done returns the channel that notifies the process quitting.
func Done() <-chan struct{} {
	return done
}

// SetWait 设置等待时间
func SetWait(duration time.Duration) {
	mutex.Lock()
	defer mutex.Unlock()
	if duration > wait {
		wait = duration
	}
}
