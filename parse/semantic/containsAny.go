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
func (c ContainsAny) IsSatisfied(val string) bool {
	for _, expected := range c.Or {
		if strings.Contains(strings.ToLower(val), strings.ToLower(expected)) {
			return true
		}
	}
	return false
}

// GetType returns routed type
func (c ContainsAny) GetType() *charlie.Type {
	return &c.Exit
}

// GetRoute Returns next route
func (c ContainsAny) GetRoute() Route {
	return c.Next
}
