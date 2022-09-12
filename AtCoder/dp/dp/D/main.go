package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func maxInList(list []int64) int64 {
	var min int64 = math.MinInt64
	for _, n := range list {
		if n < min {
			min = n
		}
	}

	return min
}

func solve(N int64, W int64, w []int64, v []int64) {
	// index 0: 品物の数
	// index 1: 品物の重さ
	// dp[index 0][index 1]: ナップサックに入っている品物の価値の総和
	var dp [][]int64
	for i := 0; i <= int(N); i++ {
		var dpRow []int64
		for j := 0; j <= int(W); j++ {
			dpRow = append(dpRow, 0)
		}
		dp = append(dp, dpRow)
	}

	var ans int64 = 0
	for i := 1; i <= int(N); i++ {
		for j := 0; j <= int(W); j++ {
			// 品物iを入れない場合
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			// 品物iを入れる場合
			diff := j - int(w[i-1])
			if 0 <= diff {
				dp[i][j] = max(dp[i][j], dp[i-1][diff]+v[i-1])
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}

	fmt.Println(ans)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	w := make([]int64, N)
	v := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, W, w, v)
}
