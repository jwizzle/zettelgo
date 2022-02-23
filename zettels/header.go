package zettels

import (
  "gopkg.in/yaml.v3"
)
// Represent the header of a note.
type Header struct {
	Title string `yaml:"title"`
	Date string `yaml:"date"`
	Tags []string `yaml:"tags"`
	Links map[string]string `yaml:"links"`
	Delimiter string
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
