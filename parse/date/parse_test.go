package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPatterns(t *testing.T) {
	for layout, r := range patterns {
		t.Run("Layout "+layout, func(t *testing.T) {
			assert.Regexp(t, r, layout)
		})
	}
}

var parseProvider = []struct {
	Data   string
	Time   time.Time
	Parsed bool
}{
	// Long year
	{"05-12-2008", date(2008, 12, 5), true},
	{"05-12-2008.", date(2008, 12, 5), true},
	{"05.12.2008", date(2008, 12, 5), true},
	{"05.12.2008,", date(2008, 12, 5), true},
	{"05/12/2008", date(2008, 12, 5), true},
	{" 05-12-2008 ", date(2008, 12, 5), true},
	{"foo 05.12.2008 23", date(2008, 12, 5), true},
	{"2006 05/12/2008 super release.", date(2008, 12, 5), true},

	// Short year
	{"05-12-08", date(2008, 12, 5), true},
	{"05.12.08", date(2008, 12, 5), true},
	{"05/12/08", date(2008, 12, 5), true},
	{" 05-12-08 ", date(2008, 12, 5), true},
	{"foo 05.12.08 23", date(2008, 12, 5), true},
	{"2006 05/12/08 super release.", date(2008, 12, 5), true},

	// False cases
	{"foo05-12-08", time.Time{}, false},
	{" 05-12-08foo", time.Time{}, false},
	{" 1205-12-200801", time.Time{}, false},

	// String month cases
	{"September 13, 2018", date(2018, 9, 13), true},
	{"Release september 13, 2018.", date(2018, 9, 13), true},
	{"Release september 13, 2018, main fixes bellow", date(2018, 9, 13), true},
	{"September 13 2018", date(2018, 9, 13), true},
	{"September 2018", date(2018, 9, 1), true},
	{"Test (September 13, 2018).", date(2018, 9, 13), true},
	{"September    13,   2018", date(2018, 9, 13), true},
}

func TestParse(t *testing.T) {
	for _, data := range parseProvider {
		t.Run("Parsing "+data.Data, func(t *testing.T) {
			res, parsed := Parse(data.Data)

			if assert.Equal(t, data.Parsed, parsed) {
				if data.Parsed {
					assert.Equal(t, res, data.Time)
				}
			}
		})
	}
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}