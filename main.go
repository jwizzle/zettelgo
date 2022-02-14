// Zettelgo is a commandline application to interface with a folder of 'zettelkast' notes.
// This main module implements the 'zettels' package to do all heavy lifting.
// Its main purpose is to handle configuration and parse CLI opts/params.
package main

import (
  "fmt"
  "os"
	"github.com/jwizzle/zettelgo/zettels"
)

var HOME, CFG_FILE string

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

// Main entry point.
func main() {
	HOME = os.Getenv("HOME")
	CFG_FILE = HOME + "/.zettelgo_conf.yaml"

	cfg := *config_init(&zettels.Cfg{
		Directory: HOME + "/.zettelkasten",
		Ignore_list: []string{
			".git",
		},
		Header_delimiter: "---",
	})

	box := zettels.Box{Config: cfg}
	_, err := box.Fill()
  if err != nil {
    panic(err)
  }

  for _, note := range box.Notes {
    fmt.Println(note.Title)
  }
}
