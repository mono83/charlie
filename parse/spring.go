package parse

import (
	"github.com/mono83/charlie"
	"regexp"
	"strings"
	"fmt"
	"time"
)

var springVersionPattern = regexp.MustCompile(`Changes in version ([\w\d.]+) \(([\w\-]+)\)`)
var springIssuePattern = regexp.MustCompile(`^\* ([\w\d\-]+) - (.*)$`)

func ParseSpringChangelog(data string) ([]charlie.ChangeLog, error) {
	var result []charlie.ChangeLog

	var current *charlie.ChangeLog
	for i, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			// Empty string

			// Storing previous
			if current != nil {
				result = append(result, *current)
			}
			current = nil
			continue
		}
		if len(line) > 2 && (line[0:2] == "--" || line[0:2] == "==") {
			// Separators
			continue
		}

		if springVersionPattern.MatchString(line) {
			chunks := springVersionPattern.FindStringSubmatch(line)

			// Parsing release
			r, ok := ParseRelease(chunks[1])
			if !ok {
				return nil, fmt.Errorf(`unable to parse release signature "%s" at line %d`, chunks[1], i+1)
			}

			// Parsing date
			t, err := time.Parse("2006-01-02", chunks[2])
			if err != nil {
				return nil, fmt.Errorf(`time parse error line %d - %s`, i+1, err.Error())
			}

			// Storing previous
			if current != nil {
				result = append(result, *current)
			}

			current = &charlie.ChangeLog{Release: *r, Date: t.UTC()}

			continue
		}

		if springIssuePattern.MatchString(line) {
			if current == nil {
				return nil, fmt.Errorf("issue signature found at line %d before release name", i+1)
			}

			chunks := springIssuePattern.FindStringSubmatch(line)

			current.Issues = append(current.Issues, charlie.Issue{
				ID:      chunks[1],
				Message: chunks[2],
			})
		}
	}

	if current != nil {
		result = append(result, *current)
	}

	return result, nil
}
