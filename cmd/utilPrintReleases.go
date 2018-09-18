package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mono83/charlie"
)

func printReleases(releases []charlie.Release) {
	for _, rel := range releases {
		color.Cyan(fmt.Sprintf("%s", rel.Version))
		for i, issue := range rel.Issues {
			fmt.Printf("%3d. ", i+1)
			colorForType(issue.Type).Print(issue.Type.String())
			fmt.Println("", issue.Message)
		}
		fmt.Println()
	}
}

func colorForType(t charlie.Type) *color.Color {
	switch t {
	case charlie.Fixed:
		return color.New(color.FgYellow)
	default:
		return color.New(color.FgGreen)
	}
}
