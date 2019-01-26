package model

import "strings"

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
	Unknown     Type = 126
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

func ParseIssueType(t string) Type {
	switch strings.ToLower(t) {
	case "info":
		return Info
	case "added":
		return Added
	case "changed":
		return Changed
	case "deprecated":
		return Deprecated
	case "fixed":
		return Fixed
	case "security":
		return Security
	case "performance":
		return Performance
	default:
		return Unknown
	}
}
