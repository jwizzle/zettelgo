package zettels

import (
	"fmt"
	"strings"

  "gopkg.in/yaml.v3"
)
// Represent the header of a note.
type Header struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Tags []string `yaml:"tags"`
	Links map[string]string `yaml:"links"`
}

// Return the header as a marshalled string.
func (self *Header) Display() (string, error) {
	ymlcont, err := yaml.Marshal(self)
	if err != nil {
		return "", err
	}
	out := fmt.Sprintf("%s", string(ymlcont))
	out = strings.ReplaceAll(out, "'", "")
	out = strings.ReplaceAll(out, "\"", "")
	return out, nil
}

// Unmarshal header bytestring to an object.
func NewHeader(text []byte, path string) (*Header, error) {
	data := Header{}

	unmarshal_err := yaml.Unmarshal(text, &data)
	if unmarshal_err != nil {
		return nil, &HeaderMalformedError{path: path}
	}
	return &data, nil
}
