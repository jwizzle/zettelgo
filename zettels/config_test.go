package zettels

import (
	"testing"
)

// Test creation, merging and instantiating from file.
func TestCfg(t *testing.T) {
	cfg := Cfg{
		Directory: tmpdir,
		Ignore_list: []string{
			".git",
		},
	}

  for _, filename := range tmpfilenames {
    if ! util.String_in_slice(tmpdir + "/" + filename, notes) {
      t.Errorf("test_box: Note not gathered: %v", filename)
    }
  }
}
