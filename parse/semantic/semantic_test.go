package semantic

import (
	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var semanticDataProvider = []struct {
	Data     string
	Route    Route
	Expected model.Type
}{
	{
		"* Add something",
		Contains{Val: "add", Exit: model.Performance},
		model.Performance,
	},
	{
		" -- FIXED something",
		Contains{Val: "fix", Exit: model.Fixed},
		model.Fixed,
	},
	{
		" -- something happened",
		Contains{
			Val:  "fix",
			Exit: model.Fixed,
			Next: &Contains{
				Val:  "add",
				Exit: model.Performance,
				Next: &AlwaysTrue{Exit: model.Info},
			},
		},
		model.Info,
	},
	{
		" -- something happened in this WORLD please fix it",
		ContainsAll{
			And:  []string{"fix", "world"},
			Exit: model.Fixed,
		},
		model.Fixed,
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
