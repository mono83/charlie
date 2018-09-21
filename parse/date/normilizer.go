package date

import "strings"

func normalize(val string) string {
	val = strings.TrimSpace(val)
	val = strings.Trim(val, "(),.")
	return val
}
