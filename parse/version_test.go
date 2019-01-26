package parse

import (
	"testing"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

var parseVersionDataProvider = []struct {
	Given    string
	Expected model.Version
}{
	{"2.0.10.RELEASE", model.Version{Major: "2", Minor: "0", Patch: "10", Label: "RELEASE"}},
	{"0.2.12", model.Version{Major: "0", Minor: "2", Patch: "12"}},
	{"3.1-beta", model.Version{Major: "3", Minor: "1", Label: "beta"}},
	{"0.1", model.Version{Major: "0", Minor: "1"}},
	{"v4.0.0-rc.1", model.Version{Major: "4", Minor: "0", Patch: "0", Label: "rc.1"}},
	{" v4.0.0-rc.1 ", model.Version{Major: "4", Minor: "0", Patch: "0", Label: "rc.1"}},
	{"Release v4.0.0-rc.1", model.Version{Major: "4", Minor: "0", Patch: "0", Label: "rc.1"}},
	{"Is it release? v4.0.0-rc.1 Mr. Charlie!", model.Version{Major: "4", Minor: "0", Patch: "0", Label: "rc.1"}},
}

func TestParseVersion(t *testing.T) {
	for _, data := range parseVersionDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			r, ok := Version(data.Given)
			if assert.True(t, ok, "unable to parse "+data.Given) {
				assert.Equal(t, data.Expected, *r)
			}
		})
	}
}
