package parse

import (
	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var issueTypeDataProvider = []struct {
	Expected     model.Type
	Given        string
	GivenHeaders []string
}{
	{model.Info, `Add a warning if React.forwardRef render function ...`, nil},
	{model.Fixed, `Add a warning if React.forwardRef render function ...`, []string{"BugFix"}},
	{model.Fixed, `Fix gridArea to be treated as a unitless CSS property (@mgol in #13550)`, nil},
	{model.Security, `Fix a potential XSS vulnerability when the attacker controls an attribute name (CVE-2018-6341).`, nil},
}

func TestIssue(t *testing.T) {
	for _, data := range issueTypeDataProvider {
		t.Run(data.Given, func(t *testing.T) {
			assert.Equal(t, data.Expected, Issue(data.Given, data.GivenHeaders).Type)
		})
	}
}
