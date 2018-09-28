package parse

import (
	"errors"
	"github.com/mono83/charlie"
	"github.com/mono83/charlie/parse/date"
	"github.com/mono83/charlie/parse/markdown"
	"github.com/mono83/charlie/parse/semantic"
	"strings"
)

var reactSemanticRoute = semantic.ContainsAny{
	Or:   []string{"fix", "bug"},
	Exit: charlie.Fixed,
	Next: &semantic.ContainsAny{
		Or:   []string{"improve", "performance"},
		Exit: charlie.Performance,
		Next: &semantic.ContainsAny{
			Or:   []string{"add", "provide"},
			Exit: charlie.Added,
			Next: &semantic.AlwaysTrue{Exit: charlie.Info},
		},
	},
}

// ReactChangelog Parses react change logs
func ReactChangelog(src charlie.Source) ([]charlie.Release, error) {
	lines := markdown.ToListLines(src.Body)

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
			if t, parsed := date.Parse(line.Headers[0]); parsed {
				lastRelease.Date = t
			}
		} else if versionDetected && *version != lastRelease.Version {
			// If new release detected while parsing lines
			releases = append(releases, *lastRelease)
			// New release
			lastRelease = &charlie.Release{Version: *version}
			if t, parsed := date.Parse(line.Headers[0]); parsed {
				lastRelease.Date = t
			}
		}

		// Detecting type
		var issueType charlie.Type
		if t, detected := semantic.Walk(reactSemanticRoute, line.Value); !detected || t == nil {
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
