package main

import (
	"github.com/mono83/charlie/drivers"
	"github.com/mono83/charlie/config/ini"
	"time"
	"github.com/mono83/charlie/process"
	"github.com/mono83/charlie/parse"
)

/*
	This is an example of different drivers/parsers configurations being attached to channel
 */
func main() {
	config, err := ini.GetDefaultConfig()
	if (err != nil) {
		panic("No configuration found")
	}

	githubDriver := drivers.GithubDriver{config.Auth.Github}
	lastProcessed := time.Unix(0, 0) // TODO need to obtain lastProcessedValue somewhere
	reactProcessor := process.GetReleaseProcessor(parse.ReactChangelog, process.PrintToConsole)
	githubDriver.ApplyToReleasesLastProcessed("facebook/react", reactProcessor, lastProcessed)
}
