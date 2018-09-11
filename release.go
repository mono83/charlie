package charlie

import "time"

type Release struct {
	Major   string
	Minor   string
	Version string
	Patch   string
	Build   string
}

type ChangeLog struct {
	Release Release
	Issues  []Issue
	Date    time.Time
}
