package drivers

import (
	"errors"
	"regexp"
	"strings"
)

// GitHubReleasesAPI is a driver, that obtains release information
// directly from GitHub releases API https://developer.github.com/v3/repos/releases/
//
// For each obtained release callback function will be invoked
// Execution will be stopped when callback returns null or there are no more releases
//
// Flow:
// 1. Driver reads latest release and sends it to callback
//    https://developer.github.com/v3/repos/releases/#get-the-latest-release
//    Example: https://api.github.com/repos/facebook/react/releases/latest
// 2. If there were no errors it fetches releases list
//    https://developer.github.com/v3/repos/releases/#list-releases-for-a-repository
// 3. TODO If there is still no error - it traverses across other pages
//
//
// PS. TODO implement authentication
func GitHubReleasesAPI(repository string, callback func(title, body string) error) error {
	if callback == nil {
		return errors.New("empty callback")
	}

	if !isValidGithubRepository(repository) {
		return errors.New("Invalid repository name")
	}

	// Reading latest release into JSON
	var rel simplifiedReleaseInfo
	err := IntoJSON(&rel)(Only200(HTTPGet("https://api.github.com/repos/" + repository + "/releases/latest")))
	if err != nil {
		return err
	}

	// Invoking callback
	if err := callback(rel.Name, rel.Body); err != nil {
		return err
	}

	// Reading list into JSON
	var list []simplifiedReleaseInfo
	err = IntoJSON(&list)(Only200(HTTPGet("https://api.github.com/repos/" + repository + "/releases")))
	if err != nil {
		return err
	}

	for _, r := range list {
		if r.Name == rel.Name {
			continue
		}

		if err := callback(r.Name, r.Body); err != nil {
			return err
		}
	}

	return nil
}

type simplifiedReleaseInfo struct {
	Name        string `json:"tag_name"`
	PublishedAt string `json:"published_at"`
	Body        string `json:"body"`
}

func isValidGithubRepository(repository string) bool {
	parts := strings.Split(repository, "/")
	if len(parts) == 2 && isValidGithubRepoName(parts[0]) && isValidGithubUserName(parts[1]) {
		return true
	}
	return false
}

var repoNameRegexp = regexp.MustCompile("(^[\\w-.]+$)")
var userNameRegexp = regexp.MustCompile("(^[\\w]([\\w-]*[\\w])?$)")
var multipleHyphenRegexp = regexp.MustCompile("-{2,}")

func isValidGithubUserName(userName string) bool {
	return userNameRegexp.MatchString(userName) && !multipleHyphenRegexp.MatchString(userName)
}

func isValidGithubRepoName(repoName string) bool {
	return repoNameRegexp.MatchString(repoName)
}
