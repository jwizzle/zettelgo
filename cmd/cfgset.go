/*
TODO
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cfgsetCmd represents the cfgset command
var cfgsetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config.",
	Long: `Key value, separated by a space. So basically your first argument is the key, second the value.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cfgset called")
	},
}

func init() {
	cfgCmd.AddCommand(cfgsetCmd)
}
