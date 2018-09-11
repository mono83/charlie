package parse

import (
	"testing"

	"github.com/mono83/charlie"
	"github.com/stretchr/testify/assert"
)

var parseVersionDataProvider = []struct {
	Given    string
	Expected charlie.Version
}{
	{"2.0.10.RELEASE", charlie.Version{Major: "2", Minor: "0", Patch: "10", Label: "RELEASE"}},
	{"0.2.12", charlie.Version{Major: "0", Minor: "2", Patch: "12"}},
	{"3.1-beta", charlie.Version{Major: "3", Minor: "1", Label: "beta"}},
	{"0.1", charlie.Version{Major: "0", Minor: "1"}},
}

func TestParseVersion(t *testing.T) {
	for _, data := range parseVersionDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			r, ok := Version(data.Given)
			if assert.True(t, ok) {
				assert.Equal(t, data.Expected, *r)
			}
		})
	}
}
