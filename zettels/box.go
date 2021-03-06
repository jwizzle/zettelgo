// Provides functions to interface with a collection of zettels/notes.
package zettels

import (
	"errors"
  "os"
	"strings"
  "path/filepath"
	"encoding/json"

	"github.com/jwizzle/zettelgo/util"
)

// Represent a collection of zettels, and provides an interface to them.
// All manipulation and interfacing with zettels should be done through the box.
type Box struct {
	Notes []Note
	Notepaths []string
	Config Cfg
}

// Return json byte representation.
func (self *Box) ToJson(filter NoteFilter) ([]byte, error) {
	var notesOut []Note

	for _, note := range self.Notes {
		if filter.Match(note, *self) {
			notesOut = append(notesOut, note)
		}
	}

	jsonbytes, err := json.Marshal(notesOut)
	if err != nil{
		return []byte{}, err
	}

	return jsonbytes, nil
}

// Return a list of all unique tags in the box.
func (self *Box) GetTags(filter NoteFilter) ([]string) {
	var uniqtags []string
	for _, note := range self.Notes {
		for _, tag := range note.Header.Tags {
			if ! util.StringInSlice(tag, uniqtags) &&
			filter.Match(note, *self){
				uniqtags = append(uniqtags, tag)
			}
		}
	}
	return uniqtags
}

// Retrieve a note from the box.
// Filtered by a json string.
func (self *Box) GetNote(filter NoteFilter) (Note, error) {
	for _, note := range self.Notes {
		if filter.MatchAny(note, *self) {
			return note, nil
		}
	}

	return Note{}, errors.New("Note not found.")
}

// Retrieve notes from the box.
// Possibly filtered by a json string.
func (self *Box) GetNotesS(filter NoteFilter) ([]Note, error) {
	notes := []Note{}
	for _, note := range self.Notes {
		if filter.Match(note, *self) {
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
		if err != nil {
			handleError(err)
		} else {
			self.Notes = append(self.Notes, newnote)
		}
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
