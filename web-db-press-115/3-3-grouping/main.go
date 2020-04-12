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
	// N=社員数は2~10
	// 社員は1~N番
	var N int

	// A=社員のペアの組み合わせに対しての幸福度
	// A[i][j]は社員iと社員jがペアの場合の幸福度
	// 幸福度は-1,000,000 以上 1,000,000 以下の整数
	var A [][]int

	N, A = Read(r)

	fmt.Println(N, A)

	//3つ以下のグループに最適に分割する組み合わせで幸福度が最大になるものを計算する

	/*
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
	*/
}

func solve(N int, A [][]int) int {
	return 0
}

func Read(r io.Reader) (int, [][]int) {
	var N int
	fmt.Fscanf(r, "%d", &N)

	return N, [][]int{
		{1, 2, 3},
		{2, 3, 4},
	}
}
