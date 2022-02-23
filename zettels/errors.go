package zettels

import (
	"fmt"
)

type HeaderMalformedError struct{path string}
func (m *HeaderMalformedError) Error() string {
	return "Skipping note with malformed header. At " + m.path
}

func handleError(e error) {
    switch e.(type) {
			case *HeaderMalformedError :
				fmt.Println(e)
			default:
				if e != nil {
						panic(e)
				}
    }
}
