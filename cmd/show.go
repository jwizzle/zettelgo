/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	showHeader bool
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of a zettel.",
	Long: `Show the contents of a zettel.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
    if len(args) != 1 {
      return errors.New("Show takes exactly 1 argument.")
    }
		splitarg := strings.Split(args[0], "/")
		pathlessArg := splitarg[len(splitarg) - 1]
		note, err := zettelBox.GetNote(pathlessArg)
		if err != nil {
			return err
		}

		var notecontent []byte
		if showHeader {
			notecontent, err = note.GetFullContent()
		} else {
			notecontent, err = note.GetContent()
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
}
