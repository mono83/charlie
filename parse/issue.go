package parse

import (
	"github.com/mono83/charlie/model"
	"strings"
)

// Issue method parses issue string with headers data
func Issue(line string, headers []string) (i model.Issue) {
	i.Message = line

	// Analyzing headers
	if len(headers) > 0 {
		for _, h := range headers {
			h = strings.ToLower(strings.TrimSpace(h))
			if strings.Contains(h, "bug") {
				i.Type = model.Fixed
			} else if strings.Contains(h, "performance improve") {
				i.Type = model.Performance
			}
		}
	}

	// Analyzing contents
	lineLower := strings.ToLower(line)
	if strings.Contains(line, "CVE") || strings.Contains(lineLower, "vulnerabilit") {
		i.Type = model.Security
	} else if strings.HasPrefix(lineLower, "fix") {
		if strings.Contains(lineLower, "performance") {
			i.Type = model.Performance
		} else {
			i.Type = model.Fixed
		}
	}

	return
}
