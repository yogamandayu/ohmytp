package route

import (
	"fmt"
	"strings"
)

// Group to group route.
func Group(method string, group string, uri string) string {
	return fmt.Sprintf("%s %s/%s", strings.ToUpper(method), group, uri)
}
