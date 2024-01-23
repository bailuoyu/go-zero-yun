package queue

type Log struct {
}

func (q Log) QueueName() string {
	return "log"
}
