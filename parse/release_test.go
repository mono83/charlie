package parse

import (
	"github.com/mono83/charlie"
	"testing"
	"github.com/stretchr/testify/assert"
)

var parseReleaseDataProvider = []struct {
	Given    string
	Expected charlie.Release
}{
	{"2.0.10.RELEASE", charlie.Release{Major: "2", Minor: "0", Patch: "10", Version: "RELEASE"}},
	{"0.2.12", charlie.Release{Major: "0", Minor: "2", Patch: "12"}},
	{"3.1-beta", charlie.Release{Major: "3", Minor: "1", Version: "beta"}},
	{"0.1", charlie.Release{Major: "0", Minor: "1"}},
}

func TestParseRelease(t *testing.T) {
	for _, data := range parseReleaseDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			r, ok := ParseRelease(data.Given)
			if assert.True(t, ok) {
				assert.Equal(t, data.Expected, *r)
			}
		})
	}
}
