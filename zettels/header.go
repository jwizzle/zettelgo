// TODO Finish
package zettels

import (
	"regexp"
)

// Represent the header of a note.
type Header struct {
	Text []string
	Sections map[string]Section
}

// Parse the header content into a list of generic sections.
// TODO Something with missing titles.
func parse_genericsections(headertext []string) ([]Section) {
	var sections []Section
	var sectiontext []string
	sectionstarted := false

	for _, line := range headertext {
		firstchar := line[0:1]

		if ! sectionstarted {
			if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(firstchar) {
				sectionstarted = true
				sectiontext = []string{line}
			}
		} else {
			if regexp.MustCompile(`^[a-zA-Z]$`).MatchString(firstchar) {
				firstline := sectiontext[0]
				sectiontitle_regex := regexp.MustCompile(`([a-zA-Z]+)`)
				newsectiontitle := sectiontitle_regex.Find([]byte(firstline))
				newsection := Section{
					Title: string(newsectiontitle),
					Contentlist: sectiontext,
				}
				sections = append(sections, newsection)
				sectiontext = []string{line}
			} else {
				sectiontext = append(sectiontext, line)
			}
		}
	}

	return sections
}

// Parse a header text into a header object.
func (self *Header) parse() (*Header, error) {
	if self.Sections == nil {
		self.Sections = make(map[string]Section)
	}

	genericsections := parse_genericsections(self.Text)
	for _, section := range genericsections {
		self.Sections[section.Title] = unpack_genericsection(section)
	}

	return self, nil
}
