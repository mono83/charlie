package semantic

import "github.com/mono83/charlie"

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
func (r AlwaysTrue) GetRoute() Route {
	return nil
}
