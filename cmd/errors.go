package cmd

import (
	"fmt"
)

type DisplayParamMalformedError struct{}
func (m *DisplayParamMalformedError) Error() string {
	return "Display param malformed."
}

func handle_error(e error) {
    switch e.(type) {
			case *DisplayParamMalformedError :
				fmt.Println(e)
			default:
				if e != nil {
						panic(e)
				}
    }
}
