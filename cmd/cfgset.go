/*
TODO
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/jwizzle/zettelgo/zettels"
	"github.com/spf13/cobra"
)

// cfgsetCmd represents the cfgset command
var cfgsetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config.",
	Long: `Takes a json string as single argument.`,
	RunE: func(cmd *cobra.Command, args []string) (error) {
		if len(args) != 1 {
			return &ArgumentError{Msg: "1 expected"}
		}

		jsonbyte := []byte(args[0])
		var newcfg zettels.Cfg
		err := json.Unmarshal(jsonbyte, &newcfg)
		if err != nil{
			return err
		}
		zettelCfg.Merge(newcfg)

		// TODO Write to file.

		fmt.Println(zettelCfg.ToString())
		return nil
	},
}

func init() {
	cfgCmd.AddCommand(cfgsetCmd)
}
