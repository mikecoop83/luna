package luna

import "fmt"

type path string

func (p path) appendKey(key string) path {
	return p + path(fmt.Sprintf("['%s']", key))
}

func (p path) appendIndex(idx int) path {
	return p + path(fmt.Sprintf("[%d]", idx))
}
