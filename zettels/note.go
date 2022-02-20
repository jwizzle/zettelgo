// TODO Finish
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
	Header Header
}

// Return the header []byte from a filepath. Without delimiters.
func headertext_from_filepath(path string, delimiter string) ([]byte, error) {
	var header []byte
	headeropened := false
	file, err := os.Open(path)
	handle_error(err)
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

// Wrap all links in the header text with '' if they aren't already.
// Needed for yaml validation. Since [[]] are invalig in unwrapped strings.
// TODO Make sure already quoted links are ignored.
func wrap_links(text []byte) ([]byte) {
	link_regexp := regexp.MustCompile(`\[\[[\w\._ ]+\]\]`)
	unquoted_links := link_regexp.FindAll(text, -1)
	for _, unq_link := range unquoted_links {
		quoted_link := []byte("\"" + string(unq_link) + "\"")
		text = bytes.ReplaceAll(text, unq_link, quoted_link)
	}

	return text
}

// Read a zettel and return a parsed Note object.
func Note_from_filepath(path string, config Cfg) (Note, error) {
	headertext, err := headertext_from_filepath(path, config.Header_delimiter)
	handle_error(err)
	headertext = wrap_links(headertext)
	newheader, err := Header_from_text(headertext)
	handle_error(err)

	return Note{
		Title: newheader.Title,
		Header: *newheader,
		Path: path,
	}, nil
}
