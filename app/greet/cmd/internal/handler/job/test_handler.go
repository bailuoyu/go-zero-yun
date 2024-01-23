package job

import (
	"github.com/spf13/cobra"
	"go-zero-yun/public/handler"
	"time"
)

// TestHandler 测试
func TestHandler(cmd *cobra.Command, args []string) {
	time.Sleep(9 * time.Second)
	handler.JobSuccess(cmd, "test job")
}
