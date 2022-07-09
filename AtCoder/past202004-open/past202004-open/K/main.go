package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, S string, C []int64, D []int64) {

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
    var S string
    scanner.Scan()
    S = scanner.Text()
    C := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    D := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        D[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, S, C, D)
}
