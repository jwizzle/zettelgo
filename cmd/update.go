/*
Copyright © 2022 jwizzle info@hossel.net
TODO finish
Random idea:
Update by json in by first building a json representation of the box like:
{
	"box": {
		*
	}
}
then allowing to update by json.
Should probably create basic functionality first tho for renaming.

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jwizzle/zettelgo/zettels"

	"github.com/spf13/cobra"
)

var (
	updateTitle string
)

func makeUpdateFilter(args []string) (zettels.NoteFilter, error) {
	var filter zettels.NoteFilter
	var err error
	if jsonFilter == "" {
		filter = zettels.NoteFilter{
			Path: args[0],
			Filename: args[0],
		}
	} else {
		filter, err = zettels.NewNoteFilter(jsonFilter)
		if err != nil {
			return zettels.NoteFilter{}, err
		}
	}

	return filter, nil
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a zettel.",
	Long: `arguments: [path/filename] or pass a custom json filter.`,
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) != 1 && jsonFilter == "" {
			return &ArgumentError{Msg: "Requires either a json filter or filepath/filename."}
    }
    return nil
  },
	RunE: func(cmd *cobra.Command, args []string) (error) {
		filter, err := makeUpdateFilter(args)
		if err != nil {
			return err
		}
		note, err := zettelBox.GetNote(filter)
		if err != nil {
			return err
		}
		newheader := zettels.Header{
			Title: note.Header.Title,
			Date: note.Header.Date,
			Tags: note.Header.Tags,
			Links: note.Header.Links,
		}

		// TODO Actually update the title of the note and write it someway
		// Create note.Move() or something.
		var newfilename string
		if updateTitle != "" {
			newheader.Title = updateTitle
			newfilename = strings.ToLower(
				strings.ReplaceAll(note.Filename, strings.ToLower(note.Header.Title), updateTitle))
		}

		headerstring, _ := newheader.Display()
		fmt.Println(headerstring)
		fmt.Println(newfilename)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&updateTitle, "title", "", "New title for the zettel.")
	filterable(updateCmd)
	jsonable(updateCmd)
}
