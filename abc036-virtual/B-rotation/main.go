package main

import (
	"fmt"
	"strings"
)

func main() {
	// 標準入力からの読み取り
	var N int
	fmt.Scanf("%d", &N)

	s := make([][]string, N)
	for i := 0; i < N; i++ {
		var line string
		fmt.Scanf("%s", &line)
		s[i] = strings.Split(line, "")
	}

	// 問題を解く
	ans := solve(s)

	// 出力する
	for i := 0; i < N; i++ {
		fmt.Println(strings.Join(ans[i], ""))
	}
}

func solve(s [][]string) [][]string {
	N := len(s)
	ss := make([][]string, N, N)
	for i := 0; i < N; i++ {
		ss[i] = make([]string, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// 1,1 => 1,N
			// 1,2 => 2,N
			// 2,3 => 3,N-1
			// i,j => j, N-1-i
			// s[i][j] は ss[j][N-1-i] に動く
			ss[j][N-1-i] = s[i][j]
		}
	}

	return ss
}
