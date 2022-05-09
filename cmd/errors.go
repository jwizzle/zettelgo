package cmd

import (
	"fmt"
)

type DisplayParamMalformedError struct{}
func (m *DisplayParamMalformedError) Error() string {
	return "Display param malformed."
}

type ArgumentError struct{Msg string}
func (m *ArgumentError) Error() string {
	return fmt.Sprintf("Argument error: %s", m.Msg)
}

func handleError(e error) {
    switch e.(type) {
			case *DisplayParamMalformedError :
				fmt.Println(e)
			default:
				if e != nil {
					fmt.Println(e)
				}
    }
}
