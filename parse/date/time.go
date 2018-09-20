package date

import (
	"fmt"
	"regexp"
	"time"
)

// ParseTime parsing string time
func ParseTime(val string) (*time.Time, bool, error) {
	// Loop
	for pattern, layout := range timePatterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return nil, false, fmt.Errorf("unable to compile pattern %s - %s", pattern, err.Error())
		}

		if !r.MatchString(val) {
			continue
		}

		match := r.FindStringSubmatch(val)

		t, err := time.Parse(layout, normalizeTime(match[0]))
		if err != nil {
			return nil, false, fmt.Errorf("time parse error by pattern %s - %s", pattern, err.Error())
		}
		return &t, true, nil
	}

	return nil, false, nil
}
