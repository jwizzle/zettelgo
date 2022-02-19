package zettels

import (
	"fmt"
	"regexp"
)

// Represent a generic section.
type Section struct {
	Title string
	Contentlist []string
}

// Represent a section with string content.
type Stringsection struct {
	Title string
	Content string
}
// Represent a section with list content.
type Listsection struct {
	Title string
	Content []string
}

// Represent a section with Links.
type Linksection struct {
	Title string
	Content []string
}

// Parse a generic section to a specific one.
// Mostly makes sure the content is of the right type.
// TODO Sub-todos, error laten returnen, en subparsers dat laten doen.
func (genericsection Section) parse() (interface{}) {
	if len(genericsection.Contentlist) == 1 {
		stringsection := Stringsection{
			Title: genericsection.Title,
			Content: genericsection.Contentlist[0],
		}
		return stringsection.parse()
	} else {
		listsection := Listsection{
			Title: genericsection.Title,
			Content: genericsection.Contentlist,
		}
		// TODO .parse maken en hier implementeren.
		listsection = listsection.parse()
		return listsection
	}
}

// Parse a stringsection, setting the content correctly.
func (self *Stringsection) parse() (Stringsection) {
	prefix_regexp := regexp.MustCompile(self.Title + ` ?[:\-\/\\\_=] ?`)
	prefix_str := string(prefix_regexp.Find([]byte(self.Content)))
	prefix_len := len(prefix_str)
	self.Content = self.Content[prefix_len:]

	return *self
}

// Parse a listsection into a tag or linksection.
// TODO Finish, figure out what to do with links and tags.
// Links now don't show up for some reason.
func (self *Listsection) parse() (Listsection) {
	for _, i := range self.Content {
		fmt.Println(i)
	}
	return *self
}
