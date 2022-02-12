package main

import (
	"testing"
)

func TestBox(t *testing.T) {
  tmpdir, tmpfilenames := CreateTmpdir(t)
	cfg := Cfg{
		Directory: tmpdir,
		Ignore_list: []string{
			".git",
		},
	}
	box := Box{Config: cfg}
	notes, err := box.gather_paths()
  if err != nil {
    panic(err)
  }

  for _, filename := range tmpfilenames {
    if ! string_in_slice(tmpdir + "/" + filename, notes) {
      t.Errorf("test_box: Note not gathered: %v", filename)
    }
  }
}

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
    result := path_in_ignorelist(scenario.path, scenario.ignorelist)

		if result != scenario.expect {
			t.Errorf(
        "test_path_in_ignorelist: path: %v, ignorelist: %v, expect %v, got: %v.",
        scenario.path, scenario.ignorelist, scenario.expect, result,
      )
		}
	}
}
