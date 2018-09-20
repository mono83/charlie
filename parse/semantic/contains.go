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
func (r Contains) IsSatisfied(val string) bool {
	return strings.Contains(strings.ToLower(val), strings.ToLower(r.Val))
}

// GetType returns routed type
func (r Contains) GetType() *charlie.Type {
	return &r.Exit
}

// GetRoute Returns next route
func (r Contains) GetRoute() Route {
	return r.Next
}
