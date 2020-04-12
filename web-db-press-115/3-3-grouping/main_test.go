package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSolve_問題文のサンプルが通ること(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{`6
10 10 -10 -10 -10
10 -10 -10 -10
-10 -10 -10
10 -10
-10`, "40"},
		{`3
1 1
1`, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			r := strings.NewReader(tt.in)
			w := &bytes.Buffer{}
			Solve(r, w)

			if got := w.String(); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
