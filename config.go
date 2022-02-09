package main

import (
  "fmt"
  "os"
  "gopkg.in/yaml.v3"
	"io/ioutil"
)

type Cfg struct {
	Directory string
	Ignore_list []string
}

func (self *Cfg) merge(other Cfg) {
	if other.Directory != "" {
		self.Directory = other.Directory
	}
	if other.Ignore_list != nil {
		for _, ignore_item := range other.Ignore_list {
			if ! string_in_slice(ignore_item, self.Ignore_list) {
				self.Ignore_list = append(self.Ignore_list, ignore_item)
			}
		}
	}
}

func cfg_from_file(path string) (*Cfg, error) {
	data := Cfg{}

	yfile, yml_err := ioutil.ReadFile(path)
	if yml_err != nil {
    switch yml_err.(type) {
			case *os.PathError :
				text := fmt.Sprintf("No config file at: %v. Continuing with default dir.",
														CFG_FILE)
				fmt.Println(text)
			default:
				return nil, yml_err
    }
	} else {
		unmarshal_err := yaml.Unmarshal(yfile, &data)
		if unmarshal_err != nil {
			return nil, unmarshal_err
		}
	}

  return &data, nil
}
