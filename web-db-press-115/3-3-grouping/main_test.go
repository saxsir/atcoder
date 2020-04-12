package main

import (
	"bytes"
	"reflect"
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

func Test入力を読み込めている(t *testing.T) {
	type W struct {
		N int
		A [][]int
	}

	tests := []struct {
		in   string
		want W
	}{
		{`6
10 10 -10 -10 -10
10 -10 -10 -10
-10 -10 -10
10 -10
-10`, W{
			N: 6, A: [][]int{
				{10, 10, -10, -10, -10},
				{10, -10, -10, -10, 0},
				{-10, -10, -10, 0, 0},
				{10, -10, 0, 0, 0},
				{-10, 0, 0, 0, 0},
			},
		}},
		{`3
1 1
1`, W{
			N: 3, A: [][]int{
				{1, 1},
				{1, 0},
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			r := strings.NewReader(tt.in)
			N, A := Read(r)

			if N != tt.want.N {
				t.Errorf("got %v, want %v", N, tt.want.N)
			}

			if reflect.DeepEqual(A, tt.want.A) != true {
				t.Errorf("got %v, want %v", A, tt.want.A)
			}
		})
	}

}
