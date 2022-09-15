package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func solve(N int64, W int64, w []int64, v []int64) {
	// index 0: item cnt
	// index 1: item value
	// weight[i][j]: item weight
	var weight [][]uint64
	for i := 0; i <= int(N); i++ {
		var weightRow []uint64
		for j := 0; j <= 100*1000; j++ {
			weightRow = append(weightRow, math.MaxInt64)
		}
		weight = append(weight, weightRow)
	}
	weight[0][0] = 0

	var ans = 0
	// 品物ループ
	for i := 1; i <= int(N); i++ {
		// 価値ループ
		for j := 0; j <= 100*1000; j++ {
			// count no item
			weight[i][j] = min(weight[i][j], weight[i-1][j])
			// count item
			diff := j - int(v[i-1])
			if 0 <= diff && 0 <= weight[i-1][diff]+uint64(w[i-1]) {
				weight[i][j] = min(weight[i][j], weight[i-1][diff]+uint64(w[i-1]))
			}

			if weight[i][j] <= uint64(W) && ans < j {
				ans = j
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
