package exchange

type DelayLog struct {
}

func (ex DelayLog) ExchangeName() string {
	return "delay_log"
}
