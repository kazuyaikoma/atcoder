package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, a []int64, u []int64, v []int64) {

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
    a := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    u := make([]int64, N-1)
    v := make([]int64, N-1)
    for i := int64(0); i < N-1; i++ {
        scanner.Scan()
        u[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, a, u, v)
}
