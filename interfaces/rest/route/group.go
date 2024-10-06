package route

import (
	"fmt"
	"strings"
)

func Group(method string, group string, uri string) string {
	return fmt.Sprintf("%s %s/%s", strings.ToUpper(method), group, uri)
}
