package parsers

import (
	"github.com/mono83/charlie"
	"github.com/mono83/charlie/parse"
)

// ParserFunc describes common parsers.
// It takes title (may be empty) and returns slice of releases
type ParserFunc func(title, body string) ([]charlie.Release, error)

// registeredParsers contains collection of registered parsers
var registeredParsers map[string]ParserFunc

// Names returns list of all registered parsers
func Names() []string {
	var names []string
	for n := range registeredParsers {
		names = append(names, n)
	}

	return names
}

// Find searches for parser by it's name
func Find(name string) (ParserFunc, bool) {
	p, ok := registeredParsers[name]
	return p, ok
}

func init() {
	registeredParsers = map[string]ParserFunc{}
	registeredParsers["spring"] = parse.SpringChangelog
	registeredParsers["react"] = parse.ReactChangelog
}
