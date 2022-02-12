package zettels

import (
  "testing"
	"os"
)

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
