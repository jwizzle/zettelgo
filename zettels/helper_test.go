package zettels

import (
  "testing"
	"os"
)

type testfile struct {
	filename string
	path string
	content string
	header_expect Header
}

// Create a temporary directory with some zettels.
// Return both the path as string and the list of present files.
func CreateTmpdir(t *testing.T) (string, []testfile) {
	tmpdir := t.TempDir()
	testfiles := []testfile{
		{
			filename: "henk.md",
			path: tmpdir + "/" + "henk.md",
			content: `
---
title: henk
date: 2022-02-01T17:19
tags:
  - #henk
  - #blaat
links:
  foo: [[foo.md]] 
---

Note content
`,
			header_expect: Header{
				Title: "henk",
				Date: "2022-02-01T17:19",
				Tags: []string{"#henk", "#blaat"},
				Links: map[string]string{
					"foo": "[[foo.md]]",
				},
			},
		},
		{
			filename: "foo.md",
			path: tmpdir + "/" + "foo.md",
			content: `
---
title: foo
date: 2022-02-01T17:19
tags:
  - #foo
  - #blaat
links:
  henk: [[henk.md]] 
---

Note content
`,
			header_expect: Header{
				Title: "foo",
				Date: "2022-02-01T17:19",
				Tags: []string{"#foo", "#blaat"},
				Links: map[string]string{
					"henk": "[[henk.md]]",
				},
			},
		},
		{
			filename: "bar.md",
			path: tmpdir + "/" + "bar.md",
			content: `
---
title: bar
date: 2022-02-01T17:19
tags:
  - #bar
  - #blaat
links:
  henk: [[henk.md]] 
---

Note content
`,
			header_expect: Header{
				Title: "bar",
				Date: "2022-02-01T17:19",
				Tags: []string{"#bar", "#blaat"},
				Links: map[string]string{
					"henk": "[[henk.md]]",
				},
			},
		},
	}

	for _, testfile := range testfiles {
			newfile, err := os.Create(testfile.path)
			if err != nil {
					panic(err)
			}
			defer newfile.Close()
			newfile.WriteString(testfile.content)
	}
	return tmpdir, testfiles
}
