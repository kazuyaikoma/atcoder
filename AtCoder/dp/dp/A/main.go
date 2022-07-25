package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int, h []int) {
	dp := []float64{0, math.Abs(float64(h[0] - h[1]))}
	for i := 2; i < N; i++ {
		diff := math.Min(
			math.Abs(float64(h[i-2]-h[i]))+dp[i-2],
			math.Abs(float64(h[i-1]-h[i]))+dp[i-1],
		)
		dp = append(dp, diff)
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
	h := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, h)
}
