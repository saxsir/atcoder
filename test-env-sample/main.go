package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Solve(os.Stdin, os.Stdout)
}

// テストから使えるように入力、出力を抽象化
func Solve(r io.Reader, w io.Writer) {
	// 入力読み込み
	var A, B int
	fmt.Fscanf(r, "%d %d", &A, &B)

	// 問題を解く
	ans := solve(A, B)

	// 答えの出力
	fmt.Fprint(w, ans)
}

// 問題を解く実装
func solve(A, B int) int {
	if B%A == 0 {
		return B / A
	}

	return (B / A) + 1
}
