// Package funckit pkg中的引用的函数包
package funckit

import (
	"fmt"
	"strconv"
	"time"
)

func DurToMic2(duration time.Duration) float64 {
	//return math.Trunc(float64(duration.Microseconds())/10+0.5) * 1e-2
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(duration.Microseconds())/1000), 64)
	return value
}
