package charlie

import (
	"testing"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

var versionCompareDataProvider = []struct {
	Expected int
	A, B     model.Version
}{
	{
		0,
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Major diff check
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "9", Minor: "2", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Minor diff check
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "1", Minor: "9", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Patch diff check
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "1", Minor: "2", Patch: "9", Build: "4", Label: "5"},
	},
	{
		-1, // Build diff check
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "9", Label: "5"},
	},
	{
		-1, // Label diff check
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "a"},
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "b"},
	},
	{
		-1, // Major diff check with different length
		model.Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		model.Version{Major: "11", Minor: "2", Patch: "3", Build: "4", Label: "5"},
	},
}

func TestVersionCompare(t *testing.T) {
	for _, data := range versionCompareDataProvider {
		t.Run(data.A.String()+" - "+data.B.String(), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.A.Compare(data.B))
			assert.Equal(t, 0-data.Expected, data.B.Compare(data.A), "Reverse check failed")
		})
	}
}

var containsDataProvider = []struct {
	Expected        bool
	Mask, Candidate model.Version
}{
	{
		true,
		model.Version{Major: "2"},
		model.Version{Major: "2", Minor: "1"},
	},
	{
		true,
		model.Version{Major: "2"},
		model.Version{Major: "2", Minor: "1", Patch: "4"},
	},
	{
		true,
		model.Version{Major: "2"},
		model.Version{Major: "2", Patch: "4"},
	},
	{
		true,
		model.Version{Major: "5", Patch: "3"},
		model.Version{Major: "5", Patch: "3"},
	},
	{
		true,
		model.Version{Major: "5", Patch: "3"},
		model.Version{Major: "5", Patch: "3", Build: "8544"},
	},
	{
		true,
		model.Version{Major: "5", Patch: "3"},
		model.Version{Major: "5", Patch: "3", Label: "alpha"},
	},
	{
		false, // No major version
		model.Version{Minor: "1"},
		model.Version{Minor: "1"},
	},
	{
		false, // Label in mask - always false
		model.Version{Major: "2", Label: "alpha"},
		model.Version{Major: "2", Minor: "1", Label: "alpha"},
	},
	{
		false, // Different majors
		model.Version{Major: "5", Patch: "3"},
		model.Version{Major: "6", Patch: "3"},
	},
	{
		false, // Different minors
		model.Version{Major: "5", Minor: "3"},
		model.Version{Major: "5", Minor: "4"},
	},
	{
		false, // Different patches
		model.Version{Major: "5", Patch: "3"},
		model.Version{Major: "5", Patch: "4"},
	},
}

func TestVersionContains(t *testing.T) {
	for _, data := range containsDataProvider {
		t.Run(data.Mask.String()+" - "+data.Candidate.String(), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Mask.Contains(data.Candidate))
		})
	}
}
