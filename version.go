package charlie

import (
	"bytes"
	"strings"
)

// Version describes only release version
type Version struct {
	Major string
	Minor string
	Patch string
	Label string
	Build string
}

// IsStable returns true if version not contains beta or
// other not stable flags
func (v Version) IsStable() bool {
	if len(v.Label) == 0 {
		return true
	}

	switch strings.ToLower(v.Label) {
	case "alpha", "beta", "rc", "rc1", "rc2", "rc3":
		return false
	}

	return true
}

// String returns string representation of version
func (v Version) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(v.Major)
	buffer.WriteString(".")
	buffer.WriteString(v.Minor)

	if len(v.Patch) != 0 {
		buffer.WriteString(".")
		buffer.WriteString(v.Patch)
	}
	if len(v.Label) != 0 {
		buffer.WriteString("-")
		buffer.WriteString(v.Label)
	}
	if len(v.Build) != 0 {
		buffer.WriteString(".")
		buffer.WriteString(v.Build)
	}

	return buffer.String()
}

// numberStringCmp is an utility method, that compares two strings,
// assuming, that both of them are numeric
func numberStringCmp(a, b string) int {
	la, lb := len(a), len(b)
	if la > lb {
		return 1
	} else if la < lb {
		return -1
	}

	return strings.Compare(a, b)
}

// Compare method compares two versions
// This method, used in sorting, will sort version in ascending order
func (v Version) Compare(other Version) int {
	// Checking majors
	if v.Major != other.Major {
		return numberStringCmp(v.Major, other.Major)
	}

	// Checking minors
	if v.Minor != other.Minor {
		return numberStringCmp(v.Minor, other.Minor)
	}

	// Checking patches
	if v.Patch != other.Patch {
		return numberStringCmp(v.Patch, other.Patch)
	}

	// Checking builds
	if v.Build != other.Build {
		return numberStringCmp(v.Build, other.Build)
	}

	// Comparing labels
	return strings.Compare(v.Label, other.Label)
}

// Contains methods returns true, if current release version can contain provided one.
// For example 5.2 can contain 5.2.1 and 5.2.12, but cannot contain 5.3 and 6.0
func (v Version) Contains(other Version) bool {
	// If current version has label - result always false
	// Or if there is no major
	if len(v.Label) > 0 || len(v.Major) == 0 {
		return false
	}

	// Checking majors - they are mandatory
	if v.Major != other.Major {
		return false
	}

	// Checking minors - it is optional
	if len(v.Minor) > 0 && v.Minor != other.Minor {
		return false
	}

	// Checking patches - it is optional
	if len(v.Patch) > 0 && v.Patch != other.Patch {
		return false
	}

	// Checking builds - it is optional
	if len(v.Build) > 0 && v.Build != other.Build {
		return false
	}

	return true
}
