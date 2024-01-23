package core

// UserInfo 用户信息
type UserInfo struct {
}

func (m UserInfo) Key(...string) string {
	return "user:info"
}

func (m UserInfo) Seconds() int {
	return 86400
}
