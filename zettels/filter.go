// TODO 
// Proper fuzzymatching
// Test
package zettels

import (
	"strings"
	"encoding/json"

	"github.com/jwizzle/zettelgo/util"
)

type NoteFilter struct {
	Title string
	Tag string
	Path string
	Filename string
}

// Check if any of the filter fields match the given note.
func (self *NoteFilter) MatchAny(note Note) (bool) {
	if self.Path != "" && self.Path == note.Path {
		return true
	}
	if self.Filename != "" && self.Filename == note.Filename {
		return true
	}
	if self.Title != "" && strings.Contains(note.Title, self.Title) {
		return true
	}
	if self.Tag != "" && (
		util.StringInSlice(self.Tag, note.Header.Tags) && 
		util.StringInSlice("#" + self.Tag, note.Header.Tags)) {
		return true
	}
	return false
}

// Check if the given note matches the filter.
func (self *NoteFilter) Match(note Note) (bool) {
	if self.Title != "" && ! strings.Contains(note.Title, self.Title) {
		return false
	}
	if self.Tag != "" && (
		! util.StringInSlice(self.Tag, note.Header.Tags) && 
		! util.StringInSlice("#" + self.Tag, note.Header.Tags)) {
		return false
	}
	if self.Path != "" && self.Path != note.Path {
		return false
	}
	if self.Filename != "" && self.Filename != note.Filename {
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
