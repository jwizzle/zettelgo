package zettels

import (
	"testing"
)

// Test header creation.
func TestNewheader(t *testing.T) {
	tmpdir, tmpfiles := CreateTmpdir(t)
	cfg := Cfg{
		Directory: tmpdir,
		Ignore_list: []string{
			".git",
		},
	}
	
	for _, tmpfile := range tmpfiles {
		headertext, err := headertextFromFilepath(tmpfile.path, cfg.Header_delimiter)
		handleError(err)
		headertext = wrapSpecialstrings(headertext)
		newheader, err := NewHeader(headertext, tmpfile.path)
		handleError(err)

		if newheader.Title != tmpfile.header_expect.Title {
			t.Errorf(
        "test_newheader: Title error filename: %v, expect: %v, got %v.",
        tmpfile.filename, tmpfile.header_expect.Title, newheader.Title,
      )
		}
		if newheader.Date != tmpfile.header_expect.Date {
			t.Errorf(
        "test_newheader: Date error filename: %v, expect: %v, got %v.",
        tmpfile.filename, tmpfile.header_expect.Date, newheader.Date,
      )
		}
		if len(newheader.Tags) != len(tmpfile.header_expect.Tags) {
			t.Errorf(
        "test_newheader: Tags error filename: %v, expect: %v, got %v.",
        tmpfile.filename, tmpfile.header_expect.Tags, newheader.Tags,
      )
		}
		if len(newheader.Links) != len(tmpfile.header_expect.Links) {
			t.Errorf(
        "test_newheader: Links error filename: %v, expect: %v, got %v.",
        tmpfile.filename, tmpfile.header_expect.Links, newheader.Links,
      )
		}
	}
}
