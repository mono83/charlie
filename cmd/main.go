package cmd

import (
	"github.com/spf13/cobra"
)

// Main is root command for charlie toolset
var Main = &cobra.Command{
	Use: "charlie",
}

func init() {
	Main.AddCommand(
		parsersList,
		parseFile,
		parseUrl,
	)
}
