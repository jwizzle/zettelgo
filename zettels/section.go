package zettels

import (
)

// Represent a generic section.
type Section struct {
	Title string
	Contentlist []string
}

// Represent a section with string content.
type Stringsection struct {
	Section
	content string
}
// Represent a section with list content.
type Listsection struct {
	Section
	content []string
}

// TODO
func unpack_genericsection(packedsection Section) (Section) {
	return packedsection
}
