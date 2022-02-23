package zettels

import (
	
)

type NoteFilter struct {
	Title string
	Tag string
}

// Check if the given note matches the filter.
func (self *NoteFilter) Match(note Note) (bool) {
	return true
}

func NoteFilterFromString(jsonstring string) (NoteFilter) {
	return NoteFilter{}
}
