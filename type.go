package charlie

// Type describes issue type
type Type byte

// List of defined issue types
const (
	Info       Type = 0
	Added      Type = 1
	Changed    Type = 2
	Deprecated Type = 3
	Removed    Type = 4
	Fixed      Type = 5
	Security   Type = 6
	Unreleased Type = 127
)
