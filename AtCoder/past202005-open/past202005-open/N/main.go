package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, Q int64, t []int64, x []int64, y []int64) {

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
    t := make([]int64, Q)
    x := make([]int64, Q)
    y := make([]int64, Q)
    for i := int64(0); i < Q; i++ {
        scanner.Scan()
        t[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, Q, t, x, y)
}
