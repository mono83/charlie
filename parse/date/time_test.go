package date

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestPatterns(t *testing.T) {
	for pattern, layout := range timePatterns {
		t.Run("Layout "+layout, func(t *testing.T) {
			assert.Regexp(t, regexp.MustCompile(pattern), layout)
		})
	}
}

var parseTimeProvider = []struct {
	Data   string
	Mask   string
	Parsed bool
}{
	// Long year
	{"05-12-2008", "02-01-2006", true},
	{"05-12-2008.", "02-01-2006", true},
	{"05.12.2008", "02.01.2006", true},
	{"05.12.2008,", "02.01.2006", true},
	{"05/12/2008", "02/01/2006", true},
	{" 05-12-2008 ", "02-01-2006", true},
	{"foo 05.12.2008 23", "02.01.2006", true},
	{"2006 05/12/2008 super release.", "02/01/2006", true},

	// Short year
	{"05-12-08", "02-01-06", true},
	{"05.12.08", "02.01.06", true},
	{"05/12/08", "02/01/06", true},
	{" 05-12-08 ", "02-01-06", true},
	{"foo 05.12.08 23", "02.01.06", true},
	{"2006 05/12/08 super release.", "02/01/06", true},

	// False cases
	{"foo05-12-08", "", false},
	{" 05-12-08foo", "", false},
	{" 1205-12-200801", "", false},

	// String month cases
	{"September 13, 2018", "January 2, 2006", true},
	{"Release september 13, 2018.", "January 2, 2006", true},
	{"Release september 13, 2018, main fixes bellow", "January 2, 2006", true},
	{"September 13 2018", "January 2, 2006", true},
	{"September 2018", "January 2, 2006", true},
	{"Test (September 13, 2018).", "January 2, 2006", true},
	{"September    13,   2018", "January 2 2006", true},
}

func TestParseTime(t *testing.T) {
	for _, data := range parseTimeProvider {
		t.Run("Parsing "+data.Data, func(t *testing.T) {
			res, parsed, err := ParseTime(data.Data)

			if assert.NoError(t, err) && assert.Equal(t, data.Parsed, parsed) {
				if data.Parsed {
					expected, _ := time.Parse(data.Mask, data.Data)
					res.Equal(expected)
				}
			}
		})
	}
}
