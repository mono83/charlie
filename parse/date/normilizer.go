package date

import "strings"

func normalizeTime(val string) string {
	val = strings.TrimSpace(val)
	val = strings.Trim(val, "(),.")
	return val
}
