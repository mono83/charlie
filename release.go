package charlie

import "time"

// Release describes whole release data - the name itself,
// release date and issues
type Release struct {
	Version
	Issues []Issue
	Date   time.Time
}
