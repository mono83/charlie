package charlie

import (
	"github.com/mono83/charlie/model"
	"sort"
)

// SortByVersion sorts provided slice of releases (by reference) in ascending order
// by version
func SortByVersion(releases []model.Release) {
	sort.Slice(releases, func(i, j int) bool {
		return releases[i].Version.Compare(releases[j].Version) < 0
	})
}

// SortByVersionDesc sorts provided slice of releases (by reference) in descending order
// by version
func SortByVersionDesc(releases []model.Release) {
	sort.Slice(releases, func(i, j int) bool {
		return releases[i].Version.Compare(releases[j].Version) > 0
	})
}
