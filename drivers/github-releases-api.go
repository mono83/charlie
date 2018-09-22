package drivers

import (
	"errors"
	"fmt"
)

// GitHubReleasesAPI is a driver, that obtains release information
// directly from GitHub releases API https://developer.github.com/v3/repos/releases/
//
// For each obtained release callback function will be invoked
// Execution will be stopped when callback returns null or there are no more releases
//
// PS. TODO implement authentication
func GitHubReleasesAPI(repository string, callback func(title, body string) error) error {
	if callback == nil {
		return errors.New("empty callback")
	}

	// TODO validate repository name

	// Reading list into JSON
	var list []simplifiedReleaseInfo
	page := 1
	for ok := true; ok; ok = len(list) > 0{
		url := fmt.Sprintf("https://api.github.com/repos/%s/releases?page=%d", repository, page)
		err := IntoJSON(&list)(Only200(HTTPGet(url)))
		if err != nil {
			return err
		}

		for _, r := range list {
			if err := callback(r.Name, r.Body); err != nil {
				return err
			}
		}

		page++
	}
	return nil
}

type simplifiedReleaseInfo struct {
	Name        string `json:"tag_name"`
	PublishedAt string `json:"published_at"`
	Body        string `json:"body"`
}
