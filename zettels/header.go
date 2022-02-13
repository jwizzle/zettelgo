// TODO Finish
package zettels

import (
	"regexp"
)

type Section struct {
	Title string
	Content string
}

type Stringsection struct {
	Section
}
type Listsection struct {
	Section
}

type Header struct {
	Text []string
	Sections []Section
}

// Parse a header text into a header object.
func (self *Header) parse() (*Header, error) {
	var sections [][]string
	var sectiontext []string
	sectionstarted := false

	// TODO Sections laten implementeren, Testen en abstraheren
	for _, line := range self.Text {
		firstchar := line[0:1]

		if ! sectionstarted {
			if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(firstchar) {
				sectionstarted = true
				sectiontext = append(sectiontext, line)
			}
		} else {
			if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(firstchar) {
				sections = append(sections, sectiontext)
				sectiontext = []string{line}
			} else {
				sectiontext = append(sectiontext, line)
			}
		}
	}

	return self, nil
}
