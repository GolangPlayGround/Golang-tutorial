package string_util

import "strings"

func MyContains(str, substr string) bool {
	return strings.Contains(str, substr)
}
