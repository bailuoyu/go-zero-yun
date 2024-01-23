package core

import "fmt"

const PreFix = "core"

func prefixName(name string) string {
	return fmt.Sprintf("%s_%s", PreFix, name)
}
