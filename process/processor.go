package process

import (
	"github.com/mono83/charlie/parse/parsers"
)

// GetReleaseProcessor returns function that contains of 2 steps:
// parsing releases from plaintext and processing releases based on provided functions for each step
func GetReleaseProcessor(parse parsers.ParserFunc, process ReleaseConsumer) func(title, body string) error {
	return func(title, body string) error {

		// Parsing releases
		releases, error := parse(title, body)
		if error != nil {
			return error
		}

		// Processing releases
		for _, release := range releases {
			if error := process(release); error != nil {
				return error
			}
		}

		return nil
	}
}
