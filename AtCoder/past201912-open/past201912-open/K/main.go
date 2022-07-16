package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int64, p []int64, Q int64, a []int64, b []int64) {

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
    p := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        p[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
    var Q int64
    scanner.Scan()
    Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    a := make([]int64, Q)
    b := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, p, Q, a, b)
}
