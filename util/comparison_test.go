package util

import (
  "testing"
)

func TestStringinslice(t *testing.T) {
  scenarios := []struct {
    text string
    slice []string
    expect bool
  }{
    {"bliep", []string{"bliep", "bloep"}, true},
    {"bloep", []string{"blaat"}, false},
  }

  for _, scenario := range scenarios {
    result := String_in_slice(scenario.text, scenario.slice)

		if result != scenario.expect {
			t.Errorf(
        "test_stringinslice: text: %v, slice: %v, expect %v, got: %v.",
        scenario.text, scenario.slice, scenario.expect, result,
      )
		}
	}
}
