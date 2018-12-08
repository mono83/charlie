package main

import (
	"fmt"
	"github.com/mono83/charlie/config/ini"
	"github.com/mono83/charlie/drivers"
	"github.com/mono83/charlie/parse"
	"github.com/mono83/charlie/process"
	"time"
)

/*
	This is an example of different drivers/parsers configurations being attached to channel
*/
func run() {

	// Creating channel for incoming requests
	requests := make(chan string)

	// Imitating requests for changelog processing
	setPeriodicRequests(requests, 5*time.Second, "react")
	setPeriodicRequests(requests, 3*time.Second, "spring")

	// Subscribing processors
	config, err := ini.GetDefaultConfig()
	if err != nil {
		panic("No configuration found")
	}

	githubDriver := drivers.GithubDriver{config.Auth.Github}
	lastProcessedTimes := make(map[string]time.Time)

	for {
		request := <-requests
		switch request {
		case "react":
			fmt.Println("### Processing request for facebook/react")
			lastProcessed, found := lastProcessedTimes[request]
			if !found {
				lastProcessed = time.Unix(0, 0)
			}
			githubDriver.ApplyToReleasesLastProcessed("facebook/react", process.GetReleaseProcessor(parse.ReactChangelog, process.PrintToConsole), lastProcessed)
			lastProcessedTimes[request] = time.Now()
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
