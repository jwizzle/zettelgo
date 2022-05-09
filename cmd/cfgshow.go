/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"reflect"

	"github.com/spf13/cobra"
)

// cfgshowCmd represents the cfgshow command
var cfgshowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all config or for a specific key.",
	Long: `Takes one or no arguments, the argument being the config key to show.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		var out string
		var err error

		if len(args) < 1 {
			out, err = zettelCfg.ToString()
			if err != nil {
				return err
			}
		} else {
			cfgreflection := reflect.ValueOf(zettelCfg)
			cfgfield := caseInsenstiveFieldByName(cfgreflection, args[0])
			if cfgfield == (reflect.Value{}) {
				out = "Unknown config field."
			} else {
				out = cfgfield.Interface().(string)
			}
		}

		if jsonOut {
			jsonbytes, err := json.Marshal(out)
			if err != nil{
				return err
			}
			fmt.Println(string(jsonbytes))
		} else {
			fmt.Println(out)
		}
		return nil
	},
}

func init() {
	cfgCmd.AddCommand(cfgshowCmd)
	jsonable(cfgshowCmd)
}
