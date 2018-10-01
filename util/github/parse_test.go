package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseStringsDataProvider = []struct {
	Given    string
	Expected Repository
}{
	{"foo/bar", Repository{User: "foo", Name: "bar"}},
	{"https://github.com/mono83/charlie", Repository{User: "mono83", Name: "charlie"}},
}

func TestParse(t *testing.T) {
	for _, data := range parseStringsDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			r, err := Parse(data.Given)
			if assert.NoError(t, err) {
				assert.Equal(t, data.Expected, *r)
			}
		})
	}
}
