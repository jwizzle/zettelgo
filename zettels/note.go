// TODO Finish
package zettels

import (
	"bufio"
	"os"
)

type Note struct {
	Title string
	Path string
	Link string
	Header string
}

// Return the header string from a filepath. Without delimiters.
// When returning an error, the returned string is the filepath.
// TODO
// - Create header struct or something
// - implement
// - test
func header_from_filepath(path string, delimiter string) (string, error) {
	var header string
	headeropened := false
	file, err := os.Open(path)
	handle_error(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ! headeropened {
			if scanner.Text() != delimiter {
				return path, &HeaderMalformedError{path: path}
			} else {
				headeropened = true
			}
		} else {
			if scanner.Text() == delimiter {
				return header, nil
			} else {
				header = header + scanner.Text()
			}
		}
	}

	return path, &HeaderMalformedError{path: path}
}

// Read a zettel and return a parsed Note object.
func Note_from_filepath(path string, config Cfg) (Note, error) {
	header, err := header_from_filepath(path, config.Header_delimiter)
	handle_error(err)

	return Note{
		Title: "Henk",
		Header: header,
		Path: path,
	}, nil
}
