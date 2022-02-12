package zettels

import (
  "os"
	"strings"
  "path/filepath"
)

// Represent a collection of zettels, and provides an interface to them.
// All manipulation and interfacing with zettels should be done through the box.
type Box struct {
	Notes []string
	Config Cfg
}

// Gather all paths of notes, using the config of the box for the root dir.
func (self *Box) Gather_paths() ([]string, error) {
  err := filepath.Walk(self.Config.Directory,
		func(path string, info os.FileInfo, err error) error {
			if ! path_in_ignorelist(path, self.Config.Ignore_list) {
				self.Notes = append(self.Notes, path)
			}
			return nil
		})

  return self.Notes, err
}

// Check if a given path is in the given ignore list. Return True/False.
func path_in_ignorelist(path string, ignore_list []string) (bool) {
	for _, ignore_item := range ignore_list {
		if strings.Contains(path, ignore_item) {
			return true
		}
	}
	return false
}
