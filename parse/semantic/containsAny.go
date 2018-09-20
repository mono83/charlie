package semantic

import (
	"github.com/mono83/charlie"
	"strings"
)

// ContainsAny must contains any of elements
type ContainsAny struct {
	Or   []string
	Exit charlie.Type
	Next Route
}

// IsSatisfied returns true if val is satisfied rules
func (r ContainsAny) IsSatisfied(val string) bool {
	for _, expected := range r.Or {
		if strings.Contains(strings.ToLower(val), strings.ToLower(expected)) {
			return true
		}
	}
	return false
}

// GetType returns routed type
func (r ContainsAny) GetType() *charlie.Type {
	return &r.Exit
}

// GetRoute Returns next route
func (r ContainsAny) GetRoute() Route {
	return r.Next
}
