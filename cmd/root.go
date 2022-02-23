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
	Long: `Interface with a folder on your system housing zettels/notes.
Explore your zettelkasten from the commandline. Relations between zettels
are mapped by placing a header at the top of your zettel in yaml format.
Notes with invalid headers are ignored. See an example below:
---
# Everything between the dashes is parsed as yaml.
title: Thread level midnight
date: 2021-11-21T12:14
tags:
  - #Awardwinningscreenplayideas
links:
  Index: [[184949_20211009_index.md]] # Links are wrapped in quotes by zettelgo.
---
Place the content of your zettel etc. here. I usually use markdown.

Config is loaded from ~/.zettelgo_conf.yaml by default. If none is present
defaults will be used. Which assumes your zettels are stored in ~/.zettelkasten.`,
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
