package markdown

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fixtureSLL = `* Zero

# This is first

* One
* Two

## This is second

* Three
* Four

#### This is third

* Five 
* Six

## This is fourth

* Seven
`

func TestToListLines(t *testing.T) {
	lines := ToListLines(fixtureSLL)

	if assert.Len(t, lines, 8) {
		assert.Equal(t, "* Zero", lines[0].Value)
		assert.Equal(t, "Zero", lines[0].String())
		assert.Equal(t, 0, lines[0].Index)
		assert.Empty(t, lines[0].Headers)

		assert.Equal(t, 5, lines[2].Index)
		assert.Equal(t, "Two", lines[2].String())
		assert.Equal(t, []string{"This is first"}, lines[2].Headers)

		assert.Equal(t, "Three", lines[3].String())
		assert.Equal(t, []string{"This is first", "This is second"}, lines[3].Headers)

		assert.Equal(t, "Six", lines[6].String())
		assert.Equal(t, []string{"This is first", "This is second", "This is third"}, lines[6].Headers)

		assert.Equal(t, 19, lines[7].Index)
		assert.Equal(t, "Seven", lines[7].String())
		assert.Equal(t, []string{"This is first", "This is fourth"}, lines[7].Headers)
	}
}
