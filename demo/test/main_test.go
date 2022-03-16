package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestNameUI(t *testing.T) {
	out, _ := nameInputUI()
	if out.Text != "Anonymous" {
		t.Errorf("Expected 'Anonymous' got '%s'", out.Text)
	}
}

func TestTypedName(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{"Name is Empty", "", "Anonymous"},
		{"Name is Jesse", "Jesse", "Jesse"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) { // subtest
			out, in := nameInputUI()
			test.Type(in, tc.input)
			if out.Text != tc.expected {
				t.Errorf("%s should give '%s' got '%s'", tc.name, tc.expected, out.Text)
			}
		})
	}
}
