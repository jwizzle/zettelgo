package zettels

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jwizzle/zettelgo/util"
	"gopkg.in/yaml.v3"
)

// Represent configuration needed to find zettels on system, and parse them correctly.
type Cfg struct {
	Directory string `yaml:"directory"`
	Ignore_list []string `yaml:"ignore_list"`
	Header_delimiter string `yaml:"header_delimiter"`
	Note_suffix string `yaml:"note_suffix"`
}

func (self *Cfg) ToString() (string, error) {
	ymlcont, err := yaml.Marshal(self)
	if err != nil {
		return "", err
	}

	return string(ymlcont), nil
}

// Merge one config with the given other.
// Overriding self with existing keys in the other config.
func (self *Cfg) Merge(other Cfg) {
	if other.Directory != "" {
		self.Directory = other.Directory
	}
	if other.Note_suffix != "" {
		self.Note_suffix = other.Note_suffix
	}
	if other.Header_delimiter != "" {
		self.Header_delimiter = other.Header_delimiter
	}
	if other.Ignore_list != nil {
		for _, ignore_item := range other.Ignore_list {
			if ! util.StringInSlice(ignore_item, self.Ignore_list) {
				self.Ignore_list = append(self.Ignore_list, ignore_item)
			}
		}
	}
}

// Load config from a yaml file. Returns an instantiated configuration.
func CfgFromFile(path string) (*Cfg, error) {
	data := Cfg{}

	yfile, yml_err := ioutil.ReadFile(path)
	if yml_err != nil {
    switch yml_err.(type) {
			case *os.PathError :
				text := fmt.Sprintf("No config file at: %v. Continuing with default dir.",
														path)
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
