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
	displayAllow []string = []string{"title", "path"}
	titleFilter string
	jsonFilter string
	tagFilter string
)

// Builds the outputstring from the display var, per note.
func buildDisplaystring(note zettels.Note) (string, error){
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
		if jsonFilter == "" {
			jsonFilter = fmt.Sprintf(`{"title": "%v", "tag": "%v"}`, titleFilter, tagFilter)
		}
		notes, err := zettelBox.GetNotesS(jsonFilter)
		handleError(err)

		for _, note := range notes {
			output, err := buildDisplaystring(note)
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
	listCmd.Flags().StringVar(&titleFilter, "title", "", "Filter results by title.")
	listCmd.Flags().StringVar(&tagFilter, "tag", "", "Filter results by tag.")
	listCmd.Flags().StringVar(&jsonFilter, "filter", "",
	`Filter by json (eg. {"title": "my little pony screenplay", "tag": "#bighitsforthefuture"}).
	If this is given, all other filterflags like "title" are ignored.`)
	listCmd.Flags().StringSliceVar(&display, "display", []string{"title"},
	`Display control. Accepts a comma separated list of:
	- title
	- path
	`)
}
