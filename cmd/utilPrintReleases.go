package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mono83/charlie/model"
)

func printReleases(releases []model.Release) {
	for _, rel := range releases {
		var name string
		if !rel.Date.IsZero() {
			name = fmt.Sprintf("%s %s", rel.Version, rel.Date.Format("Jan 02, 2006"))
		} else {
			name = fmt.Sprintf("%s", rel.Version)
		}
		color.Cyan(name)
		for i, issue := range rel.Issues {
			fmt.Printf("%3d. ", i+1)
			colorForType(issue.Type).Print(issue.Type.String())
			fmt.Println("", issue.Message)
		}
		fmt.Println()
	}
}

func colorForType(t model.Type) *color.Color {
	switch t {
	case model.Fixed:
		return color.New(color.FgYellow)
	default:
		return color.New(color.FgGreen)
	}
}
