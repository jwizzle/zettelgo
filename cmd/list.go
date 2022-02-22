/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List zettels, by default just lists the titles.",
	Long: `List all zettels in the directory found in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, note := range zettelBox.Notes {
			fmt.Println(note.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
