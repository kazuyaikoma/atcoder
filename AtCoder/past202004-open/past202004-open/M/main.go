package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(D int64, L int64, N int64, C []int64, K []int64, F []int64, T []int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var D int64
    scanner.Scan()
    D, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var L int64
    scanner.Scan()
    L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var N int64
    scanner.Scan()
    N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    C := make([]int64, D)
    for i := int64(0); i < D; i++ {
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    K := make([]int64, N)
    F := make([]int64, N)
    T := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        K[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        F[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(D, L, N, C, K, F, T)
}
