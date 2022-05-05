/*
Copyright © 2022 jwizzle info@hossel.net
TODO More than show
*/
package cmd

import (
	"fmt"
	"reflect"
	"encoding/json"

	"github.com/spf13/cobra"
)

// showcfgCmd represents the showcfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Show and set config.",
	Long: `Show and set config values for zettelgo.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		var out string
		var err error

		switch args[0] {
			case "show" :
				if len(args) < 2 {
					out, err = zettelCfg.ToString()
					if err != nil {
						return err
					}
				} else {
					cfgreflection := reflect.ValueOf(zettelCfg)
					// TODO Ugly errors
					out = cfgreflection.FieldByName(args[1]).Interface().(string)
				}
			default:
				return &DisplayParamMalformedError{}
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
