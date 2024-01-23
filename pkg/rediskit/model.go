package rediskit

type Model interface {
	Key(...string) string
	Seconds() int
}
