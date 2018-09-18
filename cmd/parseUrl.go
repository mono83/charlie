package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mono83/charlie/parse/parsers"
	"github.com/spf13/cobra"
)

var parseUrlTitle string

var parseUrl = &cobra.Command{
	Use:   "url <parser name> <file name>",
	Short: "Parses URL using requested parser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("both parser name (1st argument) and URL (2nd argument) must be provided")
		}

		// Searching for parser
		p, ok := parsers.Find(args[0])
		if !ok {
			return fmt.Errorf(`unable to find parser with name "%s"`, args[0])
		}

		// Making request
		// TODO make custom implementation of HTTP client
		// with dedicated user-agent and logging
		resp, err := http.Get(args[1])
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return fmt.Errorf("expected HTTP 200 but got %d", resp.StatusCode)
		}

		// Reading contents
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Parsing
		releases, err := p(parseUrlTitle, string(bts))
		if err != nil {
			return err
		}

		printReleases(releases)

		return nil
	},
}

func init() {
	parseUrl.Flags().StringVarP(&parseUrlTitle, "title", "t", "", "")
}
