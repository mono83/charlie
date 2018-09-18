package parse

import (
	"regexp"

	"github.com/mono83/charlie"
)

var versionParseSemver = regexp.MustCompile(`v?(\d+)\.(\d+)(.(\d+))?([.\-]([\w\d\-.]+))?`)

// Version parses release signature from string
func Version(src string) (*charlie.Version, bool) {
	if !versionParseSemver.MatchString(src) {
		return nil, false
	}

	rel := charlie.Version{}
	chunks := versionParseSemver.FindStringSubmatch(src)
	rel.Major = chunks[1]
	rel.Minor = chunks[2]
	rel.Patch = chunks[4]
	rel.Label = chunks[6]

	return &rel, true
}
