package parse

import (
	"errors"
	"github.com/mono83/charlie/model"
	"github.com/mono83/charlie/parse/date"
	"github.com/mono83/charlie/parse/markdown"
	"github.com/mono83/charlie/parse/semantic"
	"strings"
	"github.com/mono83/charlie/db"
)

var reactSemanticRoute = semantic.ContainsAny{
	Or:   []string{"fix", "bug"},
	Exit: model.Fixed,
	Next: &semantic.ContainsAny{
		Or:   []string{"improve", "performance"},
		Exit: model.Performance,
		Next: &semantic.ContainsAny{
			Or:   []string{"add", "provide"},
			Exit: model.Added,
			Next: &semantic.AlwaysTrue{Exit: model.Info},
		},
	},
}

// ReactChangelog Parses react change logs
func ReactChangelog(name, data string) ([]model.Release, error) {
	return ReactChangelogWithProjectRepo(name, data, nil)
}

// ReactChangelog Parses react change logs
func ReactChangelogWithProjectRepo(_, data string, projectRepo db.ProjectRepository) ([]model.Release, error) {
	lines := markdown.ToListLines(data)

	if len(lines) == 0 {
		return nil, errors.New("undefined markdown lines")
	}
	if len(lines[0].Headers) == 0 {
		return nil, errors.New("undefined header lines")
	}

	var releases []model.Release

	var lastRelease *model.Release
	for _, line := range lines {
		version, versionDetected := Version(line.Headers[0])

		if lastRelease == nil { // New release
			lastRelease = &model.Release{}
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
			lastRelease = &model.Release{Version: *version}
			if t, parsed := date.Parse(line.Headers[0]); parsed {
				lastRelease.Date = t
			}
		}

		// Detecting type
		var issueType model.Type
		if t, detected := semantic.Walk(reactSemanticRoute, line.Value); !detected || t == nil {
			issueType = model.Info // By default
		} else {
			issueType = *t
		}

		// Detecting component
		var components []string
		if len(line.Headers) > 1 && strings.HasPrefix(strings.ToLower(strings.TrimSpace(line.Headers[1])), "react") {
			components = append(components, line.Headers[1])
		}

		lastRelease.Issues = append(lastRelease.Issues, &model.Issue{
			Message:    line.Value,
			Type:       issueType,
			Components: components,
		})

		if projectRepo != nil {
			project, err := projectRepo.GetByName("react")
			if err != nil {
				return nil, errors.New("error during getting project `react`")
			}
			if project == nil {
				return nil, errors.New("project `react` not found")
			}
			lastRelease.ProjectID = project.ID
		}
	}

	// Last release append
	releases = append(releases, *lastRelease)

	return releases, nil
}
