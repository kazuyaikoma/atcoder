package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, K int64, L int64, p []int64, q []int64, r []int64, s []int64) {

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
    var K int64
    scanner.Scan()
    K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var L int64
    scanner.Scan()
    L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    p := make([]int64, K)
    q := make([]int64, K)
    for i := int64(0); i < K; i++ {
        scanner.Scan()
        p[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        q[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    r := make([]int64, L)
    s := make([]int64, L)
    for i := int64(0); i < L; i++ {
        scanner.Scan()
        r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        s[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, K, L, p, q, r, s)
}
