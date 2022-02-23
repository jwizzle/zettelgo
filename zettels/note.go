package zettels

import (
	"bufio"
	"os"
	"regexp"
	"bytes"
)

// Represent a note/zettel.
type Note struct {
	Title string
	Path string
	Link string
	HeaderDelimiter string
	Header Header
}

// Return the full filecontent of the note.
func (self *Note) GetFullContent() ([]byte, error) {
	content, err := os.ReadFile(self.Path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Return the content of the note, minus the header.
func (self *Note) GetContent() ([]byte, error) {
	var content []byte
	headerDelims := 0
	file, err := os.Open(self.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if headerDelims < 2 {
			if scanner.Text() == self.HeaderDelimiter {
				headerDelims = headerDelims + 1
			}
		} else {
			for _, i := range scanner.Bytes() {
				content = append(content, i)
			}
			for _, i := range []byte("\n") {
				content = append(content, i)
			}
		}
	}

	return content, nil
}

// Return the header []byte from a filepath. Without delimiters.
func headertextFromFilepath(path string, delimiter string) ([]byte, error) {
	var header []byte
	headeropened := false
	file, err := os.Open(path)
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ! headeropened {
			if scanner.Text() != delimiter {
				return nil, &HeaderMalformedError{path: path}
			} else {
				headeropened = true
			}
		} else {
			if scanner.Text() == delimiter {
				return header, nil
			} else {
				for _, i := range scanner.Bytes() {
					header = append(header, i)
				}
				for _, i := range []byte("\n") {
					header = append(header, i)
				}
			}
		}
	}

	return nil, &HeaderMalformedError{path: path}
}

// Wrap all occurences of an unquoted text in quotes.
func wrapInText(text []byte, unquoted_text []byte) ([]byte) {
	quoted := []byte("\"" + string(unquoted_text) + "\"")
	if ! bytes.Contains(text, quoted){
		text = bytes.ReplaceAll(text, unquoted_text, quoted)
	}
	return text
}

// Wrap all links and tags in the header text with '' if they aren't already.
// Needed for yaml validation. Since [[]] are invalig in unwrapped strings.
func wrapSpecialstrings(text []byte) ([]byte) {
	link_regexp := regexp.MustCompile(`\[\[[\w\._ ]+\]\]`)
	tag_regexp := regexp.MustCompile(`#\w+`)
	unquoted_links := link_regexp.FindAll(text, -1)
	unquoted_tags := tag_regexp.FindAll(text, -1)

	for _, unq_link := range unquoted_links {
		text = wrapInText(text, unq_link)
	}
	for _, unq_tag := range unquoted_tags {
		text = wrapInText(text, unq_tag)
	}

	return text
}

// Read a zettel and return a parsed Note object.
func NewNote(path string, config Cfg) (Note, error) {
	headertext, err := headertextFromFilepath(path, config.Header_delimiter)
	handleError(err)
	headertext = wrapSpecialstrings(headertext)
	newheader, err := NewHeader(headertext, path)
	handleError(err)

	return Note{
		Title: newheader.Title,
		Header: *newheader,
		Path: path,
		HeaderDelimiter: config.Header_delimiter,
	}, nil
}
