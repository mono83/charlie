package semantic

import (
	"github.com/mono83/charlie"
	"strings"
)

// Contains must contains one element
type Contains struct {
	Val  string
	Exit charlie.Type
	Next Route
}

// IsSatisfied returns true if val is satisfied rules
func (c Contains) IsSatisfied(val string) bool {
	return strings.Contains(strings.ToLower(val), strings.ToLower(c.Val))
}

// GetType returns routed type
func (c Contains) GetType() *charlie.Type {
	return &c.Exit
}

// GetRoute Returns next route
func (c Contains) GetRoute() Route {
	return c.Next
}
