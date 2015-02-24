package goinput

import "strings"

func FilterTrimString(value interface{}) interface{} {
	return strings.TrimSpace(value.(string))
}
