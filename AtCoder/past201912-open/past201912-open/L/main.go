package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, M int64, x []int64, y []int64, c []int64, X []int64, Y []int64, C []int64) {

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
    var M int64
    scanner.Scan()
    M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    x := make([]int64, N)
    y := make([]int64, N)
    c := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    X := make([]int64, M)
    Y := make([]int64, M)
    C := make([]int64, M)
    for i := int64(0); i < M; i++ {
        scanner.Scan()
        X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        Y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, M, x, y, c, X, Y, C)
}
