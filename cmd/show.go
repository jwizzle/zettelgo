/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jwizzle/zettelgo/zettels"

	"github.com/spf13/cobra"
)

var (
	showHeader bool
	headerOnly bool
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of a zettel.",
	Long: `Show the contents of a zettel.`,
  Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (error) {
		splitarg := strings.Split(args[0], "/")
		pathlessArg := splitarg[len(splitarg) - 1]
		filter := zettels.NoteFilter{
			Title: pathlessArg,
			Path: pathlessArg,
			Filename: pathlessArg,
		}
		note, err := zettelBox.GetNote(filter)
		if err != nil {
			return err
		}

		var notecontent []byte
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
		if err != nil {
			return err
		}

		fmt.Println(string(notecontent))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVar(&showHeader, "header", false, "Display header in output.")
	showCmd.Flags().BoolVar(&headerOnly, "header-only", false, "Display header only.")
}
