package cmd

import (
	"fmt"
	"sort"

	"github.com/mono83/charlie/parse/parsers"
	"github.com/spf13/cobra"
)

// parsersList command displays list of registered parsers
var parsersList = &cobra.Command{
	Use:     "parsers",
	Aliases: []string{"pl"},
	Short:   "Returns list of registered parsers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Registered parsers")
		fmt.Println("==================")

		parsers := parsers.Names()
		sort.Strings(parsers) // Using alphabetical sorting

		for i, name := range parsers {
			fmt.Printf("%3d. %s\n", i+1, name)
		}

		fmt.Println()
	},
}
