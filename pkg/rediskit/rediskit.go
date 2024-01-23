package rediskit

func GetInfo(model Model, suffix ...string) (string, int) {
	var key string
	key = model.Key(suffix...)
	return key, model.Seconds()
}
