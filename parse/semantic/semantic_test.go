package semantic

import (
	"github.com/mono83/charlie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var semanticDataProvider = []struct {
	Data     string
	Route    Route
	Expected charlie.Type
}{
	{
		"* Add something",
		Contains{Val: "add", Exit: charlie.Performance},
		charlie.Performance,
	},
	{
		" -- FIXED something",
		Contains{Val: "fix", Exit: charlie.Fixed},
		charlie.Fixed,
	},
	{
		" -- something happened",
		Contains{
			Val:  "fix",
			Exit: charlie.Fixed,
			Next: &Contains{
				Val:  "add",
				Exit: charlie.Performance,
				Next: &AlwaysTrue{Exit: charlie.Info},
			},
		},
		charlie.Info,
	},
	{
		" -- something happened in this WORLD please fix it",
		ContainsAll{
			And:  []string{"fix", "world"},
			Exit: charlie.Fixed,
		},
		charlie.Fixed,
	},
}

func TestSemantic(t *testing.T) {
	for _, data := range semanticDataProvider {
		t.Run(data.Data, func(t *testing.T) {
			res, detected := Walk(data.Route, data.Data)

			if assert.True(t, detected) {
				assert.Equal(t, data.Expected, *res)
			}
		})
	}
}
