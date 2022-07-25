package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N, K int, h []int) {
	dp := []float64{0, math.Abs(float64(h[0] - h[1]))}
	for i := 2; i < N; i++ {
		minDiff := math.Pow(10, 10)
		for j := 1; j <= K && i-j >= 0; j++ {
			diff := math.Abs(float64(h[i-j]-h[i])) + dp[i-j]
			minDiff = math.Min(minDiff, diff)
		}
		dp = append(dp, minDiff)
	}
	fmt.Println(int(dp[N-1]))
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
	var K int
	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	h := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, K, h)
}
