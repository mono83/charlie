package charlie

import "strings"

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

// IsEqual returns true if version is equally
func (v Version) IsEqual(source Version) bool {
	return v.Major == source.Major &&
		v.Minor == source.Minor &&
		v.Patch == source.Patch &&
		v.Label == source.Label &&
		v.Build == source.Build
}
