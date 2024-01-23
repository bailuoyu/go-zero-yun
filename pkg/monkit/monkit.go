package monkit

type LogMsg struct {
	Name    string
	Method  string
	Err     string
	Slow    bool
	Content string
}
