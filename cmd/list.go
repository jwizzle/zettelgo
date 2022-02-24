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
	displayAllow []string = []string{"title", "path", "filename"}
	titleFilter string
	tagFilter string
)

func makeListFilter() (zettels.NoteFilter, error) {
	var filter zettels.NoteFilter
	var err error

	if jsonFilter == "" {
		filter = zettels.NoteFilter{
			Title: titleFilter,
			Tag: tagFilter,
		}
	} else {
		filter, err = zettels.NewNoteFilter(jsonFilter)
		if err != nil {
			return zettels.NoteFilter{}, err
		}
	}

	return filter, nil
}

func makeListOut(filter zettels.NoteFilter) (string, error) {
	if jsonOut {
		jsonbytes, err := zettelBox.ToJson(filter)
		if err != nil {
			return "", err
		}

		return string(jsonbytes), nil
	} else {
		output := ""
		notes, err := zettelBox.GetNotesS(filter)
		if err != nil {
			return "", err
		}

		for _, note := range notes {
			newline, err := buildDisplaystring(note)
			if err != nil {
				return "", err
			}
			output = output + newline
		}
		return output, nil
	}
}

// Builds the outputstring from the display var, per note.
func buildDisplaystring(note zettels.Note) (string, error){
	out := ""
	for _, content := range display {
		switch content {
			case "title" :
				out = out + note.Title + " "
			case "path" :
				out = out + note.Path + " "
			case "filename" :
				out = out + note.Filename + " "
			default:
				return "", &DisplayParamMalformedError{}
		}
	}
	return out + "\n", nil
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List zettels, by default just lists the titles.",
	Long: `List all zettels in the directory found in the config file.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		filter, err := makeListFilter()
		if err != nil {
			return err
		}
		out, err := makeListOut(filter)
		if err != nil {
			return err
		}

		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringSliceVar(&display, "display", []string{"title"},
	`Display control. Accepts a comma separated list of:
	- title
	- path
	- filename
	`)
	listCmd.Flags().StringVar(&titleFilter, "title", "", "Filter results by title.")
	listCmd.Flags().StringVar(&tagFilter, "tag", "", "Filter results by tag.")
	filterable(listCmd)
	jsonable(listCmd)
}
