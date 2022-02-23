package zettels

import (
  "testing"
  "github.com/jwizzle/zettelgo/util"
)

// Test creation of a box and gathering notes.
func TestBox(t *testing.T) {
  tmpdir, tmpfiles := CreateTmpdir(t)
	cfg := Cfg{
		Directory: tmpdir,
		Ignore_list: []string{
			".git",
		},
	}
	box := Box{Config: cfg}
	notes, err := box.gatherPaths()
  if err != nil {
    panic(err)
  }

  for _, tmpfile := range tmpfiles {
    if ! util.StringInSlice(tmpfile.path, notes) {
      t.Errorf("test_box: Note not gathered: %v", tmpfile.filename)
    }
  }

  // Test if filling at least doesn't return an error.
  _, err = box.Fill()
  if err != nil {
    t.Errorf("test_box: Filling returned an error: %v", err)
  }
}

// Test detection of the ignore list from the config.
func TestPathInIgnorelist(t *testing.T) {
  scenarios := []struct {
    path string
    ignorelist []string
    expect bool
  }{
    {"/home/henk/bliep", []string{"henk", "blaat"}, true},
    {"/home/henk/bliep", []string{"blaat"}, false},
  }

  for _, scenario := range scenarios {
    result := pathInIgnorelist(scenario.path, scenario.ignorelist)

		if result != scenario.expect {
			t.Errorf(
        "test_path_in_ignorelist: path: %v, ignorelist: %v, expect %v, got: %v.",
        scenario.path, scenario.ignorelist, scenario.expect, result,
      )
		}
	}
}
