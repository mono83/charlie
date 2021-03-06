package drivers

import (
	"errors"
	"fmt"
	"github.com/mono83/charlie/http"
	"regexp"
	"strings"
	"time"
)

// GithubDriver is a driver for processing releases in Github format
type GithubDriver struct {
	Auth string
}

// ApplyToReleasesLastProcessed is similar to ApplyToReleases
// but also accepts time when this method was last called for this repository
func (d GithubDriver) ApplyToReleasesLastProcessed(repository string, callback func(title, body string) error, lastProcessed time.Time) error {
	if callback == nil {
		return errors.New("empty callback")
	}

	if !isValidGithubRepository(repository) {
		return errors.New("Invalid repository name")
	}

	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + d.Auth
	headers["If-Modified-Since"] = lastProcessed.Format("Mon, 02 Jan 2006 15:04:05 MST")

	// Reading latest release into JSON
	var rel simplifiedReleaseInfo
	url := "https://api.github.com/repos/" + repository + "/releases/latest"
	err := http.IntoJSON(&rel)(http.Only200(http.Get(http.GetParams{URL: url, Headers: headers})))
	if err != nil {
		fmt.Println("Github Driver - ", err)
		return err
	}

	// Invoking callback
	if err := callback(rel.Name, rel.Body); err != nil {
		return err
	}

	// Reading list into JSON
	var list []simplifiedReleaseInfo
	page := 1
	for ok := true; ok; ok = len(list) > 0 {
		url := fmt.Sprintf("https://api.github.com/repos/%s/releases?page=%d", repository, page)
		err := http.IntoJSON(&list)(http.Only200(http.Get(http.GetParams{URL: url, Headers: headers})))
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

		page++
	}
	return nil
}

// ApplyToReleases is a driver method, that obtains release information
// directly from GitHub releases API https://developer.github.com/v3/repos/releases/
//
// For each obtained release callback function will be invoked
// Execution will be stopped when callback returns null or there are no more releases
func (d GithubDriver) ApplyToReleases(repository string, callback func(title, body string) error) error {
	return d.ApplyToReleasesLastProcessed(repository, callback, time.Unix(0, 0))
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
