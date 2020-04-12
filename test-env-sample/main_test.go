package main

import (
	"bytes"
	"strings"
	"testing"
)

func assert(t *testing.T, in, expected string) {
	r := strings.NewReader(in)
	w := bytes.NewBufferString("")

	Solve(r, w)

	if got := w.String(); got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func Test問題文のサンプルは通ること(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"12 36", "3"},
		{"12 100", "9"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			assert(t, tt.in, tt.expected)
		})
	}
}
