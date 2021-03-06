package parse

import (
	"github.com/mono83/charlie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var issueTypeDataProvider = []struct {
	Expected     charlie.Type
	Given        string
	GivenHeaders []string
}{
	{charlie.Info, `Add a warning if React.forwardRef render function ...`, nil},
	{charlie.Fixed, `Add a warning if React.forwardRef render function ...`, []string{"BugFix"}},
	{charlie.Fixed, `Fix gridArea to be treated as a unitless CSS property (@mgol in #13550)`, nil},
	{charlie.Security, `Fix a potential XSS vulnerability when the attacker controls an attribute name (CVE-2018-6341).`, nil},
}

func TestIssue(t *testing.T) {
	for _, data := range issueTypeDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			assert.Equal(t, data.Expected, Issue(data.Given, data.GivenHeaders).Type)
		})
	}
}
