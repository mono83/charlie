package github

import (
	"errors"
	"regexp"
)

var fullStringParseRegex = regexp.MustCompile(`^(https:\/\/github\.com\/)?(\w+)\/(\w+)$`)

// Parse method parses given string into GitHub repository struct
func Parse(str string) (*Repository, error) {
	if !fullStringParseRegex.MatchString(str) {
		return nil, errors.New("provided string is not GitHub repository")
	}

	matches := fullStringParseRegex.FindStringSubmatch(str)

	repo := Repository{User: matches[2], Name: matches[3]}
	if err := repo.Validate(); err != nil {
		return nil, err
	}

	return &repo, nil
}
