package semantic

import (
	"github.com/mono83/charlie/model"
)

// Route detecting type of line
type Route interface {
	// IsSatisfied returns true if val is satisfied rules
	IsSatisfied(val string) bool

	// GetType returns routed type
	GetType() *model.Type

	// GetRoute Returns next route
	GetRoute() Route
}

// Walk walks by semantic rules and returns type
func Walk(route Route, val string) (*model.Type, bool) {
	if route == nil {
		return nil, false
	}

	if route.IsSatisfied(val) {
		if t := route.GetType(); t != nil {
			return t, true
		}
		return nil, false
	}

	return Walk(route.GetRoute(), val)
}
