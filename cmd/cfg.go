/*
Copyright © 2022 jwizzle info@hossel.net
*/
package cmd

import (
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

func caseInsenstiveFieldByName(v reflect.Value, name string) reflect.Value {
    name = strings.ToLower(name)
    return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
}

// showcfgCmd represents the showcfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Show and set config.",
	Long: `Show and set config values for zettelgo.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		return &ArgumentError{Msg: "Argument expected"}
	},
}

func init() {
	rootCmd.AddCommand(cfgCmd)
	jsonable(cfgCmd)
}
