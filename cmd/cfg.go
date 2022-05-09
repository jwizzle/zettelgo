/*
Copyright © 2022 jwizzle info@hossel.net
TODO More than show
*/
package cmd

import (
	"encoding/json"
	"fmt"
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
		var out string
		var err error

		if len(args) < 1 {
			return &ArgumentError{}
		}

		switch args[0] {
			case "show" :
				if len(args) < 2 {
					out, err = zettelCfg.ToString()
					if err != nil {
						return err
					}
				} else {
					cfgreflection := reflect.ValueOf(zettelCfg)
					cfgfield := caseInsenstiveFieldByName(cfgreflection, args[1])
					if cfgfield == (reflect.Value{}) {
						out = "Unknown config field."
					} else {
						out = cfgfield.Interface().(string)
					}
				}
			default:
				return &ArgumentError{}
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
	rootCmd.AddCommand(cfgCmd)
	jsonable(cfgCmd)
}
