package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, Q int64, xmin []int64, ymin []int64, D []int64, C []int64, A []int64, B []int64) {

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
    var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    xmin := make([]int64, N)
    ymin := make([]int64, N)
    D := make([]int64, N)
    C := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        xmin[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        ymin[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    A := make([]int64, Q)
    B := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, Q, xmin, ymin, D, C, A, B)
}
