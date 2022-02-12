package zettels

import (
  "testing"
	"os"
)

// Create a temporary directory with some zettels.
// Return both the path as string and the list of present files.
func CreateTmpdir(t *testing.T) (string, []string) {
	tmpdir := t.TempDir()
	filenames := []string{"henk.md", "bliep.md", "bloep.md", "blaat.md"}

	for _, file := range filenames {
			newfile, err := os.Create(tmpdir + "/" + file)
			if err != nil {
					panic(err)
			}
			defer newfile.Close()
	}
	return tmpdir, filenames
}
