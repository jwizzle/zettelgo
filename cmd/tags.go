/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all tags.",
	Long: `List all tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, tag := range zettelBox.GetTags() {
			fmt.Println(tag)
		}
	},
}

func init() {
	listCmd.AddCommand(tagsCmd)
}
