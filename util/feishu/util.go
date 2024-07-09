package feishu

import "strings"

func ReplaceAtTag(value string) string {
	value = strings.ReplaceAll(value, "\\u003c", "<")
	value = strings.ReplaceAll(value, "\\u003e", ">")
	return value
}
