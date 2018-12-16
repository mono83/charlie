package main

import (
	"fmt"
	"github.com/mono83/charlie/config/ini"
	"github.com/mono83/charlie/drivers"
	"github.com/mono83/charlie/http"
	"github.com/mono83/charlie/parse"
	"github.com/mono83/charlie/process"
	"strings"
	"time"
)

/*
	This is an example of different drivers/parsers configurations being attached to channel
*/
func run() {

	// Creating channel for incoming requests
	requests := make(chan string)

	// Imitating requests for changelog processing
	setPeriodicRequests(requests, 50*time.Second, "react")
	setPeriodicRequests(requests, 5*time.Second, "spring-data/jpa")
	setPeriodicRequests(requests, 12*time.Second, "spring-data/commons")

	// Subscribing processors
	config, err := ini.GetDefaultConfig()
	if err != nil {
		panic("No configuration found")
	}

	githubDriver := drivers.GithubDriver{Auth: config.Auth.Github}
	lastProcessedTimes := make(map[string]time.Time)

	for {
		request := <-requests
		fmt.Println("### Processing request for", request)
		switch {
		case request == "react":
			lastProcessed, found := lastProcessedTimes[request]
			if !found {
				lastProcessed = time.Unix(0, 0)
			}
			githubDriver.ApplyToReleasesLastProcessed("facebook/react", process.GetReleaseProcessor(parse.ReactChangelog, process.PrintToConsole), lastProcessed)
			lastProcessedTimes[request] = time.Now()
		case strings.HasPrefix(request, "spring"):
			bytes, err := http.ReadFile(parse.SpringChangeLogFileURL(request))
			if err != nil {
				panic("Could not read release file for " + request)
			}
			process.GetReleaseProcessor(parse.SpringChangelog, process.PrintToConsole)("", string(bytes))
		default:
			fmt.Println("### Processing for", request, "is not implemented yet")
		}
	}
}

func setPeriodicRequests(requests chan string, period time.Duration, request string) {
	go func() {
		ticker := time.NewTicker(period)
		for range ticker.C {
			requests <- request
		}
	}()
}
