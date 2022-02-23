package zettels

import (
	"errors"
  "os"
	"strings"
  "path/filepath"
)

// Represent a collection of zettels, and provides an interface to them.
// All manipulation and interfacing with zettels should be done through the box.
type Box struct {
	Notes []Note
	Notepaths []string
	Config Cfg
}

// Retrieve a note from the box.
func (self *Box) GetNote(arg string) (Note, error) {
	for _, note := range self.Notes {
		if strings.Contains(note.Title, arg) || strings.Contains(note.Path, arg) {
			return note, nil
		}
	}
	return Note{}, errors.New("Note not found.")
}

// Retrieve notes from the box.
// Possibly filtered by a json string.
func (self *Box) GetNotesS(filterstring string) ([]Note, error) {
	notes := []Note{}
	filter, err := NewNoteFilter(filterstring)
	if err != nil {
		return []Note{}, err
	}

	for _, note := range self.Notes {
		if filter.Match(note) {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// Fill up the box with notes gathered from disk.
// Returns a reference to the filled box.
func (self *Box) Fill() (*Box, error) {
	paths, err := self.gatherPaths()
	handleError(err)

	for _, path := range paths {
		newnote, err := NewNote(path, self.Config)
		handleError(err)
		self.Notes = append(self.Notes, newnote)
	}

	return self, nil
}

// Gather all paths of notes, using the config of the box for the root dir.
func (self *Box) gatherPaths() ([]string, error) {
  err := filepath.Walk(self.Config.Directory,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if ! pathInIgnorelist(path, self.Config.Ignore_list) {
				self.Notepaths = append(self.Notepaths, path)
			}
			return nil
		})

  return self.Notepaths, err
}

// Check if a given path is in the given ignore list. Return True/False.
func pathInIgnorelist(path string, ignore_list []string) (bool) {
	for _, ignore_item := range ignore_list {
		if strings.Contains(path, ignore_item) {
			return true
		}
	}
	return false
}
