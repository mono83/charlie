package charlie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var versionCompareDataProvider = []struct {
	Expected int
	A, B     Version
}{
	{
		0,
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Major diff check
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "9", Minor: "2", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Minor diff check
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "1", Minor: "9", Patch: "3", Build: "4", Label: "5"},
	},
	{
		-1, // Patch diff check
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "1", Minor: "2", Patch: "9", Build: "4", Label: "5"},
	},
	{
		-1, // Build diff check
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "1", Minor: "2", Patch: "3", Build: "9", Label: "5"},
	},
	{
		-1, // Label diff check
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "a"},
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "b"},
	},
	{
		-1, // Major diff check with different length
		Version{Major: "1", Minor: "2", Patch: "3", Build: "4", Label: "5"},
		Version{Major: "11", Minor: "2", Patch: "3", Build: "4", Label: "5"},
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
	Mask, Candidate Version
}{
	{
		true,
		Version{Major: "2"},
		Version{Major: "2", Minor: "1"},
	},
	{
		true,
		Version{Major: "2"},
		Version{Major: "2", Minor: "1", Patch: "4"},
	},
	{
		true,
		Version{Major: "2"},
		Version{Major: "2", Patch: "4"},
	},
	{
		true,
		Version{Major: "5", Patch: "3"},
		Version{Major: "5", Patch: "3"},
	},
	{
		true,
		Version{Major: "5", Patch: "3"},
		Version{Major: "5", Patch: "3", Build: "8544"},
	},
	{
		true,
		Version{Major: "5", Patch: "3"},
		Version{Major: "5", Patch: "3", Label: "alpha"},
	},
	{
		false, // No major version
		Version{Minor: "1"},
		Version{Minor: "1"},
	},
	{
		false, // Label in mask - always false
		Version{Major: "2", Label: "alpha"},
		Version{Major: "2", Minor: "1", Label: "alpha"},
	},
	{
		false, // Different majors
		Version{Major: "5", Patch: "3"},
		Version{Major: "6", Patch: "3"},
	},
	{
		false, // Different minors
		Version{Major: "5", Minor: "3"},
		Version{Major: "5", Minor: "4"},
	},
	{
		false, // Different patches
		Version{Major: "5", Patch: "3"},
		Version{Major: "5", Patch: "4"},
	},
}

func TestVersionContains(t *testing.T) {
	for _, data := range containsDataProvider {
		t.Run(data.Mask.String()+" - "+data.Candidate.String(), func(t *testing.T) {
			assert.Equal(t, data.Expected, data.Mask.Contains(data.Candidate))
		})
	}
}
