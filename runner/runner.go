package main

import (
	"fmt"
	"github.com/mono83/charlie/config/ini"
	"github.com/mono83/charlie/config"
	"github.com/mono83/charlie/drivers"
	"github.com/mono83/charlie/http"
	"github.com/mono83/charlie/parse"
	"github.com/mono83/charlie/process"
	"strings"
	"time"
	"github.com/mono83/charlie/db/mysql"
)

/*
	This is an example of different drivers/parsers configurations being attached to channel
*/
func main() {

	// Creating channel for incoming requests
	requests := make(chan string)

	// Imitating requests for changelog processing
	setPeriodicRequests(requests, 5*time.Second, "react")
	setPeriodicRequests(requests, 50*time.Second, "spring-data/jpa")
	setPeriodicRequests(requests, 12*time.Second, "spring-data/commons")

	// Subscribing processors
	conf, err := ini.GetDefaultConfig()
	if err != nil {
		panic("No configuration found")
	}

	githubDriver := drivers.GithubDriver{Auth: conf.Auth.Github}
	lastProcessedTimes := make(map[string]time.Time)

	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		panic("Error during getting DB connection")
	}

	releaseRepo := mysql.NewMysqlReleaseRepository(db)

	for {
		request := <-requests
		switch {
		case request == "react":
			lastProcessed, found := lastProcessedTimes[request]
			if !found {
				lastProcessed = time.Unix(0, 0)
			}
			err := githubDriver.ApplyToReleasesLastProcessed("facebook/react", process.GetReleaseProcessor(parse.ReactChangelog, process.DbSaver(releaseRepo)), lastProcessed)
			if err != nil {
				fmt.Println(err)
				break
			}
			lastProcessedTimes[request] = time.Now()
		case strings.HasPrefix(request, "spring"):
			bytes, err := http.ReadFile(parse.SpringChangeLogFileURL(request))
			if err != nil {
				panic(fmt.Sprintf("Could not read release file for %s, error - %s", request, err))
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
