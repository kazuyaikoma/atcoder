package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func listMax(l []int64) int64 {
	var ret int64 = math.MinInt64
	for _, n := range l {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func solve(N int64, W int64, w []int64, v []int64) {
	var dp [][]int64
	for i := 0; i <= int(N); i++ {
		var ws []int64
		for j := 0; j <= int(W); j++ {
			ws = append(ws, math.MinInt64)
		}
		dp = append(dp, ws)
	}
	dp[0][0] = 0

	var ans int64 = 0
	for i := 1; i <= int(N); i++ {
		for j := 0; j <= int(W); j++ {
			// 品物iを入れない場合
			dp[i][j] = max(dp[i][j], dp[i-1][j])

			// 品物iを入れる場合
			// j: 袋の容量(重さ)
			// w[i]: 新たに入れる品物の重さ
			wDiff := j - int(w[i-1])
			if wDiff >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][wDiff]+v[i-1])
			}
		}
	}
	ans = listMax(dp[N])
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
