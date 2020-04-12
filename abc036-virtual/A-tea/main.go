package main

import "fmt"

func main() {
	// 標準入力からの読み取り
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	// 問題を解く
	ans := solve(A, B)
	fmt.Println(ans)
}

func solve(A, B int) int {
	if B%A == 0 {
		return B / A
	}

	return (B / A) + 1
}
