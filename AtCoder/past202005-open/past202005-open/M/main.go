package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, M int64, u []int64, v []int64, s int64, K int64, t []int64) {

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
    u := make([]int64, M)
    v := make([]int64, M)
    for i := int64(0); i < M; i++ {
        scanner.Scan()
        u[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    var s int64
    scanner.Scan()
    s, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var K int64
    scanner.Scan()
    K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    t := make([]int64, K)
    for i := int64(0); i < K; i++ {
        scanner.Scan()
        t[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, M, u, v, s, K, t)
}
