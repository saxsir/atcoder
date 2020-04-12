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
	// 入力
	var N int
	fmt.Fscanf(r, "%d", &N)

	A := make([][]int, N-1)
	for i, _ := range A {
		A[i] = make([]int, N-i-1)
	}

	// だめだ。。。また今度
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-2; j++ {
			var a int
			fmt.Scanf("%d", &a)
			A[i][j] = a
		}
	}

	// 計算
	ans := solve(N, A)

	// 出力
	fmt.Fprint(w, ans)
}

func solve(N int, A [][]int) int {
	return 0
}

func readInput(r io.Reader) {
	var N int
	fmt.Fscanf(r, "%d", &N)
}
