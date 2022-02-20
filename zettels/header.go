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
}

// Unmarshal header text to an object.
// TODO Figure out way to make sure any header parse failures
// Result in a warning.
func NewHeader(text []byte) (*Header, error) {
	data := Header{}

	unmarshal_err := yaml.Unmarshal(text, &data)
	if unmarshal_err != nil {
		return nil, unmarshal_err
	}
	return &data, nil
}
