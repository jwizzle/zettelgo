/*
Copyright © 2022 jwizzle info@hossel.net

*/
package cmd

import (
	"os"
  "fmt"

	"github.com/jwizzle/zettelgo/zettels"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zettelgo",
	Short: "A zettelkasten management tool written in Go.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		handleError(err)
		os.Exit(1)
	}
}

var (
	// Used for flags.
	userHome string
	cfgPath string
	zettelCfg zettels.Cfg
	zettelBox zettels.Box
)

func init() {
	// Persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "config file (default is $HOME/.zettelgo_conf.yaml)")

	// Handle merging of config etc.
	postinit()
}

// Instantiate a new config, by combining the defaults that are hardcoded
// and those read from '~/.zettelgo_conf.yaml' and CLI opts.
func configInit(defaults *zettels.Cfg) (*zettels.Cfg) {
	userCfg, err := zettels.CfgFromFile(cfgPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defaults.Merge(*userCfg)

	return defaults
}

// Post-init config handling.
func postinit() {
	if cfgPath == ""{
		userHome = os.Getenv("HOME")
		cfgPath = userHome + "/.zettelgo_conf.yaml"
	}

	zettelCfg = *configInit(&zettels.Cfg{
		Directory: userHome + "/.zettelkasten",
		Ignore_list: []string{
			".git",
		},
		Header_delimiter: "---",
	})
	zettelBox = zettels.Box{Config: zettelCfg}
	_, err := zettelBox.Fill()
  if err != nil {
    panic(err)
  }
}
