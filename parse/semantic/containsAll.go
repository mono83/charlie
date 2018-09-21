package semantic

import (
	"github.com/mono83/charlie"
	"strings"
)

// ContainsAll must contains all elements
type ContainsAll struct {
	And  []string
	Exit charlie.Type
	Next Route
}

// IsSatisfied returns true if val is satisfied rules
func (c ContainsAll) IsSatisfied(val string) bool {
	for _, expected := range c.And {
		if !strings.Contains(strings.ToLower(val), strings.ToLower(expected)) {
			return false
		}
	}
	return true
}

// GetType returns routed type
func (c ContainsAll) GetType() *charlie.Type {
	return &c.Exit
}

// GetRoute Returns next route
func (c ContainsAll) GetRoute() Route {
	return c.Next
}
