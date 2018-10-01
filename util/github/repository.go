package github

import (
	"errors"
	"fmt"
	"regexp"
)

// Repository contains information about github repository
type Repository struct {
	User, Name string
}

func (r Repository) String() string {
	return r.User + "/" + r.Name
}

// GetAPIURLReleasesLatest return URL to obtain latest release data
// using GitHub releases API https://developer.github.com/v3/repos/releases/
func (r Repository) GetAPIURLReleasesLatest() string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/releases/latest",
		r.User,
		r.Name,
	)
}

// GetAPIURLReleasesPage returns URL to fetch page of release data
func (r Repository) GetAPIURLReleasesPage(page int) string {
	if page < 1 {
		page = 1
	}

	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/releases?page=%d",
		r.User,
		r.Name,
		page,
	)
}

var repoNameRegexp = regexp.MustCompile("(^[\\w-.]+$)")
var userNameRegexp = regexp.MustCompile("(^[\\w]([\\w-]*[\\w])?$)")
var multipleHyphenRegexp = regexp.MustCompile("-{2,}")

// Validate validates repository name and owner
func (r Repository) Validate() error {
	if !userNameRegexp.MatchString(r.User) {
		return fmt.Errorf(`invalid repository owner user name "%s"`, r.User)
	}
	if !repoNameRegexp.MatchString(r.Name) {
		return fmt.Errorf(`invalid repository name "%s"`, r.Name)
	}
	if multipleHyphenRegexp.MatchString(r.User) {
		return errors.New("multiple hyphens found in owner user name")
	}
	if multipleHyphenRegexp.MatchString(r.Name) {
		return errors.New("multiple hyphens found in repository name")
	}

	return nil
}
