package model

import "time"

// Release describes whole release data - the name itself,
// release date and issues
type Release struct {
	ID        int64
	ProjectID int64
	Version
	Issues    []*Issue
	Date      time.Time
}

// SummaryType returns count issues by type
func (r Release) SummaryType() map[Type]int {
	summary := make(map[Type]int)

	for _, issue := range r.Issues {
		if _, ok := summary[issue.Type]; ok {
			summary[issue.Type]++
		} else {
			summary[issue.Type] = 1
		}
	}

	return summary
}
