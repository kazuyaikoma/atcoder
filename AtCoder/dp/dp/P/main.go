package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64, x []int64, y []int64) {

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
    x := make([]int64, N-1)
    y := make([]int64, N-1)
    for i := int64(0); i < N-1; i++ {
        scanner.Scan()
        x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, x, y)
}