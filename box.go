package main

import (
  "os"
	"strings"
  "path/filepath"
)

type Box struct {
	Notes []string
	Config Cfg
}

func (self *Box) gather_paths() ([]string, error) {
  err := filepath.Walk(self.Config.Directory,
		func(path string, info os.FileInfo, err error) error {
			if ! path_in_ignorelist(path, self.Config.Ignore_list) {
				self.Notes = append(self.Notes, path)
			}
			return nil
		})

  return self.Notes, err
}

func path_in_ignorelist(path string, ignore_list []string) (bool) {
	for _, ignore_item := range ignore_list {
		if strings.Contains(path, ignore_item) {
			return true
		}
	}
	return false
}
