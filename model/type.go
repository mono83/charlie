package model

// Type describes issue type
type Type byte

// List of defined issue types
const (
	Info        Type = 0
	Added       Type = 1
	Changed     Type = 2
	Deprecated  Type = 3
	Removed     Type = 4
	Fixed       Type = 5
	Security    Type = 6
	Performance Type = 7
	Unreleased  Type = 127
)

// String returns string representation of issue type
func (t Type) String() string {
	switch t {
	case Info:
		return "Info"
	case Added:
		return "Added"
	case Changed:
		return "Changed"
	case Deprecated:
		return "Deprecated"
	case Fixed:
		return "Fixed"
	case Security:
		return "Security"
	case Performance:
		return "Performance"
	default:
		return "Unknown"
	}
}
