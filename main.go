package main

import (
  "fmt"
  "os"
	"github.com/jwizzle/zettelgo/zettels"
)

var HOME, CFG_FILE string

func config_init(defaults *zettels.Cfg) (*zettels.Cfg) {
	user_cfg, err := zettels.Cfg_from_file(CFG_FILE)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defaults.Merge(*user_cfg)

	return defaults
}

func main() {
	HOME = os.Getenv("HOME")
	CFG_FILE = HOME + "/.zettelgo_conf.yaml"

	cfg := *config_init(&zettels.Cfg{
		Directory: HOME + "/.zettelkasten",
		Ignore_list: []string{
			".git",
		},
	})

	box := zettels.Box{Config: cfg}
	_, err := box.Gather_paths()
  if err != nil {
    panic(err)
  }

  for _, file := range box.Notes {
    fmt.Println(file)
  }
}
