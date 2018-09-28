package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/mono83/charlie"
	"github.com/mono83/charlie/parse/parsers"
	"github.com/spf13/cobra"
)

var parseFileTitle string

var parseFile = &cobra.Command{
	Use:   "file <parser name> <file name>",
	Short: "Parses file using requested parser",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("both parser name (1st argument) and file name (2nd argument) must be provided")
		}

		// Searching for parser
		p, ok := parsers.Find(args[0])
		if !ok {
			return fmt.Errorf(`unable to find parser with name "%s"`, args[0])
		}

		// Reading file contents
		bts, err := ioutil.ReadFile(args[1])
		if err != nil {
			return err
		}

		// Parsing
		releases, err := p(charlie.Source{Title: parseFileTitle, Body: string(bts)})
		if err != nil {
			return err
		}

		printReleases(releases)

		return nil
	},
}

func init() {
	parseFile.Flags().StringVarP(&parseFileTitle, "title", "t", "", "")
}
