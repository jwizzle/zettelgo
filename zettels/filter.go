// TODO
// Proper fuzzymatching
// Test
package zettels

import (
	"encoding/json"
	"strings"

	"github.com/jwizzle/zettelgo/util"
)

type NoteFilter struct {
	Title string
	Tag string
	Path string
	Filename string
	Link string
	LinkedFrom string `json:"linked_from"` // Matches if a link to the matching note is present in note under filename.
	LinkingTo string `json:"linking_to"` // Matches if a link to the matching note is present in note under filename.
}

// Check if any of the filter fields match the given note.
func (self *NoteFilter) MatchAny(note Note, zettelbox Box) (bool) {
	if self.Path != "" && self.Path == note.Path {
		return true
	}
	if self.Filename != "" && self.Filename == note.Filename {
		return true
	}
	if self.Title != "" && strings.Contains(note.Title, self.Title) {
		return true
	}
	if self.Link != "" {
		strippedlink := strings.ReplaceAll(self.Link, "[", "")
		strippedlink = strings.ReplaceAll(strippedlink, "]", "")
		subfilter := NoteFilter{
			Title: strippedlink,
			Path: strippedlink,
			Filename: strippedlink,
		}
		if subfilter.MatchAny(note, zettelbox) {
			return true
		}
	}
	if self.LinkedFrom != "" {
		linkedfromNote, err := zettelbox.GetNote(NoteFilter{Filename: self.LinkedFrom})
		// TODO Handle error
		if err == nil {
			if linkedfromNote.HasLink(note.Path, zettelbox) {
				return true
			}
		}
	}
	if self.LinkingTo != "" {
		linkingtoNote, err := zettelbox.GetNote(NoteFilter{Filename: self.LinkingTo})
		// TODO Handle error
		if err == nil {
			if note.HasLink(linkingtoNote.Path, zettelbox) {
				return true
			}
		}
	}
	if self.Tag != "" && (
		util.StringInSlice(self.Tag, note.Header.Tags) && 
		util.StringInSlice("#" + self.Tag, note.Header.Tags)) {
		return true
	}
	return false
}

// Check if the given note matches the filter.
func (self *NoteFilter) Match(note Note, zettelbox Box) (bool) {
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
	if self.Link != "" {
		strippedlink := strings.ReplaceAll(self.Link, "[", "")
		strippedlink = strings.ReplaceAll(strippedlink, "]", "")
		subfilter := NoteFilter{
			Title: strippedlink,
			Path: strippedlink,
			Filename: strippedlink,
		}
		if ! subfilter.MatchAny(note, zettelbox) {
			return false
		}
	}
	if self.LinkedFrom != "" {
		linkedfromNote, err := zettelbox.GetNote(NoteFilter{Filename: self.LinkedFrom})
		// TODO Handle error
		if err == nil {
			if ! linkedfromNote.HasLink(note.Path, zettelbox) {
				return false
			}
		}
	}
	if self.LinkingTo != "" {
		linkingtoNote, err := zettelbox.GetNote(NoteFilter{Filename: self.LinkingTo})
		// TODO Handle error
		if err == nil {
			if ! note.HasLink(linkingtoNote.Path, zettelbox) {
				return false
			}
		}
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
