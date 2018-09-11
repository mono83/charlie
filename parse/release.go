package parse

import (
	"github.com/mono83/charlie"
	"regexp"
)

var releaseParseSemver = regexp.MustCompile(`^(\d+)\.(\d+)(.(\d+))?([.\-]([\w\d]+))?$`)

// ParseRelease parses release signature from string
func ParseRelease(src string) (*charlie.Release, bool) {
	if !releaseParseSemver.MatchString(src) {
		return nil, false
	}

	rel := charlie.Release{}
	chunks := releaseParseSemver.FindStringSubmatch(src)
	rel.Major = chunks[1]
	rel.Minor = chunks[2]
	rel.Patch = chunks[4]
	rel.Version = chunks[6]

	return &rel, true
}
