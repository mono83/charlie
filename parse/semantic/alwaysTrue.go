package semantic

import "github.com/mono83/charlie"

// AlwaysTrue always true
type AlwaysTrue struct {
	Exit charlie.Type
}

// IsSatisfied returns true if val is satisfied rules
func (AlwaysTrue) IsSatisfied(val string) bool {
	return true
}

// GetType returns routed type
func (a AlwaysTrue) GetType() *charlie.Type {
	return &a.Exit
}

// GetRoute Returns next route
func (AlwaysTrue) GetRoute() Route {
	return nil
}
