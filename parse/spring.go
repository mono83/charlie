package parse

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mono83/charlie"
)

var springVersionPattern = regexp.MustCompile(`Changes in version ([\w\d.]+) \(([\w\-]+)\)`)
var springIssuePattern = regexp.MustCompile(`^\* ([\w\d\-]+) - (.*)$`)

// SpringChangelog parses Spring framework changelog
func SpringChangelog(_, data string) ([]charlie.Release, error) {
	var result []charlie.Release

	var current *charlie.Release
	for i, line := range strings.Split(data, "\n") {
		line = Trim(line)

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
			v, ok := Version(chunks[1])
			if !ok {
				return nil, fmt.Errorf(`unable to parse release signature "%s" at line %d`, chunks[1], i+1)
			}

			// Parsing date
			t, err := time.Parse("2006-01-2", chunks[2])
			if err != nil {
				return nil, fmt.Errorf(`time parse error line %d - %s`, i+1, err.Error())
			}

			// Storing previous
			if current != nil {
				result = append(result, *current)
			}

			current = &charlie.Release{Version: *v, Date: t.UTC()}

			continue
		}

		if springIssuePattern.MatchString(line) {
			if current == nil {
				return nil, fmt.Errorf("issue signature found at line %d before release name", i+1)
			}

			chunks := springIssuePattern.FindStringSubmatch(line)

			issue := Issue(chunks[2], nil)
			issue.ID = chunks[1]

			current.Issues = append(current.Issues, issue)
		}
	}

	if current != nil {
		result = append(result, *current)
	}

	return result, nil
}

// SpringChangeLogFileURL builds url for spring project changelog file
func SpringChangeLogFileURL(projectName string) string {
	return "https://docs.spring.io/" + projectName + "/docs/current/changelog.txt"
}
