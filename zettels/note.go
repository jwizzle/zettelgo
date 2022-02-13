// TODO Finish
package zettels

import (
	"bufio"
	"os"
)

type Note struct {
	Title string
	Path string
	Date string
	Link string
	sections []string
	Tags []string
	Links []string
	Header string
}

// Return the header string from a filepath. Without delimiters.
// When returning an error, the returned string is the filepath.
// TODO
// - implement
// - test
func header_from_filepath(path string) (string, error) {
	var header string
	headeropened := false
	file, err := os.Open(path)
	handle_error(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ! headeropened {
			// TODO Make header delimiter configurable
			if scanner.Text() != "---" {
				return path, &HeaderMalformedError{path: path}
			} else {
				headeropened = true
			}
		} else {
			if scanner.Text() == "---" {
				return header, nil
			} else {
				header = header + scanner.Text()
			}
		}
	}

	// TODO Handle better
	return path, &HeaderMalformedError{path: path}
}

// Read a zettel and return a parsed Note object.
// TODO Fill the rest of the elements
// Possibly seperate section struct, maybe look at inline struct definition for that.
func Note_from_filepath(path string) (Note, error) {
	header, err := header_from_filepath(path)
	handle_error(err)

	return Note{
		Title: "Henk",
		Header: header,
	}, nil
}
