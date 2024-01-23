package joblgc

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// SignalLogic 测试
func SignalLogic(cmd *cobra.Command, _ []string) {
	for {
		select {
		case <-cmd.Context().Done():
			fmt.Printf("ctx done time: %s\n", time.Now().Format(time.RFC3339))
			return
		case <-time.After(3 * time.Second):
			fmt.Printf("signal time: %s\n", time.Now().Format(time.RFC3339))
		}
	}
}
