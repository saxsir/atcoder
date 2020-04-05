package main

import (
	"fmt"
)

func main() {
	// 入力
	var N int
	fmt.Scanf("%d", &N)

	A := make([][]int, N-1)
	for i, _ := range A {
		A[i] = make([]int, N-i-1)
	}

	fmt.Println(A)

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
	fmt.Println(ans)
}

func solve(N int, A [][]int) int {
	fmt.Println(N)
	fmt.Println(A)
	return 0
}
