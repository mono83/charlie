package cmd

import (
	"errors"
	"fmt"
	"github.com/mono83/charlie/drivers"
	"github.com/mono83/charlie/parse/parsers"
	"github.com/spf13/cobra"
)

var parseURLTitle string

var parseURL = &cobra.Command{
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
		code, bts, err := drivers.HTTPGet(args[1])
		if err != nil {
			return err
		}
		if code != 200 {
			return fmt.Errorf("expected HTTP 200 but got %d", code)
		}

		// Parsing
		releases, err := p(parseURLTitle, string(bts))
		if err != nil {
			return err
		}

		printReleases(releases)

		return nil
	},
}

func init() {
	parseURL.Flags().StringVarP(&parseURLTitle, "title", "t", "", "")
}
