package zettels

import (
	"os"
	"log"
)

type HeaderMalformedError struct{path string}
func (m *HeaderMalformedError) Error() string {
	return "Skipping note with malformed header. At " + m.path
}

func handleError(e error) {
		l := log.New(os.Stderr, "", 0)

    switch e.(type) {
			case *HeaderMalformedError :
				l.Println(e)
			default:
				if e != nil {
						panic(e)
				}
    }
}
