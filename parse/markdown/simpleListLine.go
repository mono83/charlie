package markdown

import "strings"

// SimpleListLine contains information about exactly one
// line from Markdown document with type list (<ol> or <ul>)
type SimpleListLine struct {
	Index   int      // Contains line number
	Value   string   // Contains line contents
	Headers []string // Contains headers
}

// String returns line value without Markdown prefix
// This is also Stringer interface implementation
func (s SimpleListLine) String() string {
	return strings.TrimSpace(s.Value[1:])
}

// ToListLines parses Markdown document and returns list
// of lines
func ToListLines(src string) []SimpleListLine {
	var arr []SimpleListLine
	var hlPrevious int

	headers := make([]string, 6)
	for i, line := range strings.Split(src, "\n") {
		line := strings.TrimSpace(line)

		if strings.HasPrefix(line, "* ") {
			// This is line
			arr = append(arr, SimpleListLine{
				Index:   i,
				Value:   line,
				Headers: copyFill(headers),
			})
		} else if strings.HasPrefix(line, "#") {
			// This is header
			hl := headerSize(line)
			if hl < len(headers) {
				headers[hl-1] = strings.TrimSpace(line[hl:])
			}
			if hl < hlPrevious {
				for i := hl; i <= hlPrevious; i++ {
					headers[i] = ""
				}
			}
			hlPrevious = hl
		}
	}

	return arr
}

// copyFill copies data from one slice to another,
// ignoring empty values
func copyFill(src []string) []string {
	var arr []string
	for _, e := range src {
		if len(e) > 0 {
			arr = append(arr, e)
		}
	}
	return arr
}

// headerSize calculates Markdown header
func headerSize(line string) (size int) {
	for j, k := range line {
		if k != '#' {
			break
		}
		size = j + 1
	}
	return
}
