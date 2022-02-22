/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"

	"github.com/jwizzle/zettelgo/zettels"

	"github.com/spf13/cobra"
)

var (
	display []string
	display_allow []string = []string{"title", "path"}
)

// Builds the outputstring from the display var, per note.
func build_outputstring(note zettels.Note) (string, error){
	out := ""

	for _, content := range display {
		switch content {
			case "title" :
				out = out + note.Title + " "
			case "path" :
				out = out + note.Path + " "
			default:
				return "", &DisplayParamMalformedError{}
		}
	}

	return out, nil
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List zettels, by default just lists the titles.",
	Long: `List all zettels in the directory found in the config file.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		for _, note := range zettelBox.Notes {
			output, err := build_outputstring(note)
			if err != nil {
				return err
			}
			fmt.Println(output)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringSliceVar(&display, "display", []string{"title"},
	`Display control. Accepts a comma separated list of:
	- title
	- path
	`)
}
