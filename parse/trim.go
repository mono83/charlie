package parse

import (
	"regexp"
	"strings"
)

var dedupSpaces = regexp.MustCompile(`\s{2,}`)

// Trim method removes trailing and leading spaces
// Also this method joins two or more consecutive spaces into one
func Trim(line string) string {
	if len(line) > 0 {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			line = dedupSpaces.ReplaceAllString(line, " ")
		}
	}

	return line
}
