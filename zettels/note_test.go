package zettels

import (
	"testing"
)

// Test parsing of a note from filepath.
func TestNotefromfile(t *testing.T) {
	tmpdir, tmpfiles := CreateTmpdir(t)
	cfg := Cfg{
		Directory: tmpdir,
		Ignore_list: []string{
			".git",
		},
	}

	for _, tmpfile := range tmpfiles {
		newnote, _ := Note_from_filepath(tmpfile.path, cfg)
		expect_note := Note{
			Title: tmpfile.header_expect.Title,
			Path: tmpfile.path,
			Header: tmpfile.header_expect,
		}
		if newnote.Title != expect_note.Title{
      t.Errorf("test_note_from_file: Note title parse error: %v", tmpfile.filename)
		}
		if newnote.Path != expect_note.Path{
      t.Errorf("test_note_from_file: Note path parse error: %v", tmpfile.filename)
		}
		if newnote.Header.Title != expect_note.Header.Title{
      t.Errorf("test_note_from_file: Note header parse error: %v", tmpfile.filename)
		}
	}
}

// Test wrapping of links.
func TestLinkwrapping(t *testing.T) {
  scenarios := []struct {
    bytein []byte
    expect []byte
  }{
    {[]byte("blaatert [[henk]] ding"), []byte(`blaatert "[[henk]]" ding`)},
    {[]byte(`blaatert "[[henk]]" ding`), []byte(`blaatert "[[henk]]" ding`)},
  }

  for _, scenario := range scenarios {
    result := wrap_specialstrings(scenario.bytein)

		if string(result) != string(scenario.expect) {
			t.Errorf(
        "test_linkwrapping: bytein: %v, expect: %v, got %v.",
        string(scenario.bytein), string(scenario.expect), string(result),
      )
		}
	}
}
