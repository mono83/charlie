package parse

import (
	"fmt"
	"github.com/mono83/charlie"
	"github.com/mono83/charlie/parse/markdown"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"time"
)

var reactSemanticRoute = ContainsAny{
	Or:   []string{"fix", "bug"},
	Exit: charlie.Fixed,
	Next: &ContainsAny{
		Or:   []string{"add", "improve", "provide"},
		Exit: charlie.Performance,
		Next: &AlwaysTrue{Exit: charlie.Info},
	},
}

var reactReleaseDatePattern = regexp.MustCompile(`\(([\w\d, ]+)\)`)

// ReactChangelog Parses react change logs
func ReactChangelog(data string) ([]charlie.Release, error) {
	lines := markdown.ToListLines(data)

	if len(lines) == 0 {
		return nil, errors.New("undefined markdown lines")
	}
	if len(lines[0].Headers) == 0 {
		return nil, errors.New("undefined header lines")
	}

	var releases []charlie.Release

	var lastRelease *charlie.Release
	for _, line := range lines {
		version, versionDetected := Version(line.Headers[0])

		if lastRelease == nil { // New release
			lastRelease = &charlie.Release{}
			if versionDetected {
				lastRelease.Version = *version
			}
			if t, err := detectTime(line.Headers[0]); err == nil {
				lastRelease.Date = *t
			}
		} else if versionDetected && !version.IsEqual(lastRelease.Version) {
			// If new release detected while parsing lines
			releases = append(releases, *lastRelease)
			// New release
			lastRelease = &charlie.Release{Version: *version}
			if t, err := detectTime(line.Headers[0]); err == nil {
				lastRelease.Date = *t
			}
		}

		// Detecting type
		var issueType charlie.Type
		if t, err := SemanticWalk(reactSemanticRoute, line.Value); err != nil || t == nil {
			issueType = charlie.Info // By default
		} else {
			issueType = *t
		}

		// Detecting component
		var components []string
		if len(line.Headers) > 1 && strings.HasPrefix(strings.ToLower(strings.TrimSpace(line.Headers[1])), "react") {
			components = append(components, line.Headers[1])
		}

		lastRelease.Issues = append(lastRelease.Issues, charlie.Issue{
			Message:    line.Value,
			Type:       issueType,
			Components: components,
		})
	}

	// Last release append
	releases = append(releases, *lastRelease)

	return releases, nil
}

func detectTime(val string) (*time.Time, error) {
	if reactReleaseDatePattern.MatchString(val) {
		chunks := reactReleaseDatePattern.FindStringSubmatch(val)
		t, err := time.Parse("January 2, 2006", chunks[1])
		if err != nil {
			return nil, fmt.Errorf("time parse error - %s", err.Error())
		}
		return &t, nil
	}

	return nil, errors.New("undefined time pattern")
}
