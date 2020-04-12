package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Solve(os.Stdin, os.Stdout)
}

func Solve(r io.Reader, w io.Writer) {
	in := readInput(r)
	ans := solve(in)

	fmt.Fprint(w, ans)
}

// 問題を解く実装
func solve(s string) string {
	return s
}

func readInput(r io.Reader) string {
	var s string
	fmt.Fscanf(r, "%s", &s)

	return s
}
