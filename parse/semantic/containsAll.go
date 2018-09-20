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
func (r ContainsAll) IsSatisfied(val string) bool {
	for _, expected := range r.And {
		if !strings.Contains(strings.ToLower(val), strings.ToLower(expected)) {
			return false
		}
	}
	return true
}

// GetType returns routed type
func (r ContainsAll) GetType() *charlie.Type {
	return &r.Exit
}

// GetRoute Returns next route
func (r ContainsAll) GetRoute() Route {
	return r.Next
}
