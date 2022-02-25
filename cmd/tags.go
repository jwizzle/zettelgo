/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all tags.",
	Long: `List all tags.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		filter, err := makeListFilter()
		if err != nil {
			return err
		}
		out := ""
		if jsonOut {
			jsonout, err := json.Marshal(zettelBox.GetTags(filter))
			if err != nil {
				return err
			}
			out = string(jsonout)
		} else {
			for _, tag := range zettelBox.GetTags(filter) {
				out = out + tag + "\n"		
			}
		}

		fmt.Println(out)
		return nil
	},
}

func init() {
	listCmd.AddCommand(tagsCmd)
	jsonable(tagsCmd)
	filterable(tagsCmd)
}
