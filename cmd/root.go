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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Instantiate a new config, by combining the defaults that are hardcoded
// and those read from '~/.zettelgo_conf.yaml' and CLI opts.
func config_init(defaults *zettels.Cfg) (*zettels.Cfg) {
	user_cfg, err := zettels.Cfg_from_file(CFG_FILE)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defaults.Merge(*user_cfg)

	return defaults
}

var (
	// Used for flags.
	HOME string
	CFG_FILE string
	zettelCfg zettels.Cfg
	zettelBox zettels.Box
)

func init() {
	// Persistent flags
	rootCmd.PersistentFlags().StringVar(&CFG_FILE, "config", "", "config file (default is $HOME/.zettelgo_conf.yaml)")

	// Handle merging of config etc.
	post_init()
}

// Post-init config handling.
func post_init() {
	if CFG_FILE == ""{
		HOME = os.Getenv("HOME")
		CFG_FILE = HOME + "/.zettelgo_conf.yaml"
	}

	zettelCfg = *config_init(&zettels.Cfg{
		Directory: HOME + "/.zettelkasten",
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
