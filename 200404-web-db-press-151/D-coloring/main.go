package main

import (
	"fmt"
)

const mod int64 = 1e9 + 7

func main() {
	// 標準入力からの読み取り
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	b := make([]int, N)
	// iは1~ = 島の番号が1~
	for i := 1; i < N; i++ {
		fmt.Scanf("%d %d", &a[i], &b[i])
	}

	// 問題を解く
	ans := solve(N, a, b)

	// 出力する
	fmt.Println(ans & MOD)
}

func solve(N int, a, b []int) int {
	// 島が5個ある
	// 橋は 2-5, 1-5, 2-4, 3-2 でつながっている
	// 両端の島が黒で、、というのは1つの橋を対象にということ（経由した先とかではない）よね...
	// この場合、
	// 1が黒だったら、5は白、2は黒、3は白、4は白 => 1通り
	// 1が白だったら、5は黒か白
	// 5が黒なら、2は白... => 4通り
	// 5が白なら、2は白か黒
	// 2が白なら => 4通り
	// 2が黒なら => 2通り
	// 11通り..
	// 全部白とかもあるか

	// すべての島は必ず経由していけることが分かっているので、仮に1をスタートとして考えてみる
	// なので、1が白の場合と黒の場合に分けて考える
	// 入力からツリー構造をつくって順にたどっていく？

	return 0
}

// f(2) = f(5) * f(4) * f(3) + g(5) * g(4) * g(3)
// f(x) = g(x) + g(y1) * g(y2) * ... * g(yk)
// k=

func dfs(x, p int) {
}

func f(x, p int) {
}

func g(x, p int) {
	if dpF[x] > 0 {
		return dpF[x]
	}

	w := g(x, p)
	b := 1

}
