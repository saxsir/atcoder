package main

import "fmt"

func main() {
	// 入力
	var N int
	fmt.Scanf("%d", &N)

	// 計算
	ans := solve(N)

	// 出力
	fmt.Println(ans)
}

func solve(N int) string {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == N {
				return "Yes"
			}
		}
	}

	return "No"
}
