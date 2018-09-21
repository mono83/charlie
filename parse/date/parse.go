package date

import (
	"time"
)

// Parse parsing string date
func Parse(val string) (time.Time, bool) {
	// Loop
	for layout, r := range patterns {
		if !r.MatchString(val) {
			continue
		}

		match := r.FindStringSubmatch(val)

		if t, err := time.Parse(layout, normalize(match[0])); err == nil {
			return t, true
		}
	}

	return time.Time{}, false
}
