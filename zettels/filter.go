// TODO 
// Test
// Fuzzymatching
package zettels

import (
	"encoding/json"
	"github.com/jwizzle/zettelgo/util"
)

type NoteFilter struct {
	Title string
	Tag string
}

// Check if the given note matches the filter.
func (self *NoteFilter) Match(note Note) (bool) {
	if self.Title != "" && self.Title != note.Title {
		return false
	}
	if self.Tag != "" && ! util.StringInSlice(self.Tag, note.Header.Tags) {
		return false
	}
	return true
}

// Create a notefilter from a json string.
func NewNoteFilter(jsonstring string) (NoteFilter, error) {
	if jsonstring == "" {
		return NoteFilter{}, nil
	}
	jsonbyte := []byte(jsonstring)
	var noteFilter NoteFilter
	err := json.Unmarshal(jsonbyte, &noteFilter)
	if err != nil{
		return NoteFilter{}, err
	}
	return noteFilter, nil
}
