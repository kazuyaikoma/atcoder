package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, Q int, ss [][]int) {
	// ユーザーのフォロー状況を反映した隣接行列
	// 隣接リストも考えたが、フォロワーの把握をするには隣接行列の方が扱いやすいこと・空間計算量の制約に余裕があることから隣接行列を選択。
	adjMap := map[int][]bool{}
	// 隣接行列の初期化
	for i := 1; i <= N; i++ {
		adjMap[i] = []bool{}
		for j := 0; j <= N; j++ {
			adjMap[i] = append(adjMap[i], false)
		}
	}

	// クエリ処理
	for _, s := range ss {
		switch s[0] {
		case 1:
			adjMap[s[1]][s[2]] = true
		case 2:
			for i := 1; i <= N; i++ {
				if adjMap[i][s[1]] {
					adjMap[s[1]][i] = true
				}
			}
		case 3:
			// フォローフォローの任意の実行が次の実行に影響を与えないよう、targetMapで後で変更を適用するようにしている
			targetMap := map[int][]int{}
			for i := 1; i <= N; i++ {
				for j := 1; j <= N; j++ {
					if adjMap[s[1]][i] && adjMap[i][j] && s[1] != i && s[1] != j {
						if _, ok := targetMap[s[1]]; !ok {
							targetMap[s[1]] = []int{}
						}
						targetMap[s[1]] = append(targetMap[s[1]], j)
					}
				}
			}
			for k, vs := range targetMap {
				for _, v := range vs {
					adjMap[k][v] = true
				}
			}
		}
	}

	// 出力
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if adjMap[i][j] {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		fmt.Println("")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())

	var ss [][]int
	for i := 0; i < Q; i++ {
		scanner.Scan()
		s1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s2, _ := strconv.Atoi(scanner.Text())
		s := []int{s1, s2}
		if s1 == 1 {
			scanner.Scan()
			s3, _ := strconv.Atoi(scanner.Text())
			s = append(s, s3)
		}
		ss = append(ss, s)
	}

	solve(N, Q, ss)
}
