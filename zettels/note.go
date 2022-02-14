// TODO Finish
package zettels

import (
	"bufio"
	"os"
)

// Represent a note/zettel.
type Note struct {
	Title string
	Path string
	Link string
	Header Header
}

// Return the header string from a filepath. Without delimiters.
func headertext_from_filepath(path string, delimiter string) ([]string, error) {
	var header []string
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
				header = append(header, scanner.Text())
			}
		}
	}

	return nil, &HeaderMalformedError{path: path}
}

// Read a zettel and return a parsed Note object.
func Note_from_filepath(path string, config Cfg) (Note, error) {
	headertext, err := headertext_from_filepath(path, config.Header_delimiter)
	handle_error(err)
	newheader := Header{Text: headertext}
	_, parse_err := newheader.parse()
	handle_error(parse_err)

	// TODO
	return Note{
		Title: newheader.Sections["title"].Contentlist[0],
		Header: newheader,
		Path: path,
	}, nil
}
