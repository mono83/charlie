package parse

import (
	"github.com/mono83/charlie"
	"github.com/pkg/errors"
	"strings"
)

// SemanticRoute detecting type of line
type SemanticRoute interface {
	// IsSatisfied returns true if val is satisfied rules
	IsSatisfied(val string) bool

	// GetType returns routed type
	GetType() *charlie.Type

	// GetRoute Returns next route
	GetRoute() SemanticRoute
}

// SemanticWalk walks by semantic rules and returns type
func SemanticWalk(route SemanticRoute, val string) (*charlie.Type, error) {
	if route == nil {
		return nil, errors.New("Nil semantic route")
	}

	if route.IsSatisfied(val) {
		return route.GetType(), nil
	}

	return SemanticWalk(route.GetRoute(), val)
}

// AlwaysTrue always true
type AlwaysTrue struct {
	Exit charlie.Type
}

// IsSatisfied returns true if val is satisfied rules
func (r AlwaysTrue) IsSatisfied(val string) bool {
	return true
}

// GetType returns routed type
func (r AlwaysTrue) GetType() *charlie.Type {
	return &r.Exit
}

// GetRoute Returns next route
func (r AlwaysTrue) GetRoute() SemanticRoute {
	return nil
}

// ContainsAny must contains any of elements
type ContainsAny struct {
	Or   []string
	Exit charlie.Type
	Next SemanticRoute
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
func (r ContainsAny) GetRoute() SemanticRoute {
	return r.Next
}

// ContainsAll must contains all elements
type ContainsAll struct {
	And  []string
	Exit charlie.Type
	Next SemanticRoute
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
func (r ContainsAll) GetRoute() SemanticRoute {
	return r.Next
}

// Contains must contains one element
type Contains struct {
	Val  string
	Exit charlie.Type
	Next SemanticRoute
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
func (r Contains) GetRoute() SemanticRoute {
	return r.Next
}
