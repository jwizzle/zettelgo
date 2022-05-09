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
	showHeader bool
	headerOnly bool
)

func makeShowFilter(args []string) (zettels.NoteFilter, error) {
	var filter zettels.NoteFilter
	var err error
	if jsonFilter == "" {
		filter = zettels.NoteFilter{
			Title: args[0],
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

func makeShowOut(note zettels.Note) ([]byte, error) {
	var notecontent []byte
	var err error
	if jsonOut {
		notecontent, err = note.ToJson()
	} else {
		if headerOnly {
			var headerstring string
			headerstring, err = note.Header.Display()
			notecontent = []byte(headerstring)
		} else {
			if showHeader {
				notecontent, err = note.GetFullContent()
			} else {
				notecontent, err = note.GetContent()
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return notecontent, nil
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of a zettel.",
	Long: `Show the contents of a zettel. By default the matching can be done quite
loosely depending on your input. A full path is always matched first, and should
be consistent. A filename should be precise, it tries to match these second.
When using a title to filter as the argument. The first note where the argument is
a substring of the title of that note, that note is returned.
This might yield unexpected results. For example if you search for a note with the
title "her" but a note with the title "where I keep the bodies" is evaluated first.`,
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) != 1 && jsonFilter == "" {
			return &ArgumentError{Msg: "Requires either a json filter or argument."}
    }
    return nil
  },
	RunE: func(cmd *cobra.Command, args []string) (error) {
		filter, err := makeShowFilter(args)
		if err != nil {
			return err
		}
		note, err := zettelBox.GetNote(filter)
		if err != nil {
			return err
		}
		out, err := makeShowOut(note)
		if err != nil {
			return err
		}

		fmt.Println(string(out))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVar(&showHeader, "header", false, "Display header in output.")
	showCmd.Flags().BoolVar(&headerOnly, "header-only", false, "Display header only.")
	filterable(showCmd)
	jsonable(showCmd)
}
