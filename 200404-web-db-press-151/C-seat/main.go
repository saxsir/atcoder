package main

import (
	"fmt"
	"sort"
)

func main() {
	// 標準入力からの読み取り
	var N int
	fmt.Scanf("%d", &N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// 問題を解く
	ans := solve(a)

	// 出力する
	for i := 0; i < N; i++ {
		fmt.Println(ans[i])
	}
}

func solve(a []int) []int {
	N := len(a)
	b := make([]int, N)

	// どう解くか...出てくる数値の最大公約数で割ってみる？
	// 1は0...割合は保持しなくてもいいんだよなぁ
	// 搭乗する数値の数によって変わりそうか
	// 2種類なら2値(0, 1)でいい
	// なので、登場する数値の種類を調べ、変換表をつくる
	// 変換して出力する

	// 変換表の作成
	u := uniq(a)
	sort.Ints(u)
	m := make(map[int]int)
	for i := 0; i < len(u); i++ {
		m[u[i]] = i
	}

	// 変換
	for i := 0; i < N; i++ {
		b[i] = m[a[i]]
	}

	return b
}

func uniq(s []int) []int {
	k := make(map[int]bool)
	u := make([]int, 0, len(s))
	for _, n := range s {
		if _, v := k[n]; !v {
			k[n] = true
			u = append(u, n)
		}
	}

	return u
}
