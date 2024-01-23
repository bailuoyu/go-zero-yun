package core

type Log struct {
}

func (tp Log) TopicName() string {
	return "demo_log"
}
