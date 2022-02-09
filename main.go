package main

import (
  "fmt"
  "os"
)

var HOME, CFG_FILE string

func config_init(defaults *Cfg) (*Cfg) {
	user_cfg, err := cfg_from_file(CFG_FILE)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defaults.merge(*user_cfg)

	return defaults
}

func main() {
	HOME = os.Getenv("HOME")
	CFG_FILE = HOME + "/.zettelgo_conf.yaml"

	cfg := *config_init(&Cfg{
		Directory: HOME + "/.zettelkasten",
		Ignore_list: []string{
			".git",
		},
	})

	box := Box{Config: cfg}
	_, err := box.gather_paths()
  if err != nil {
    panic(err)
  }

  for _, file := range box.Notes {
    fmt.Println(file)
  }
}
