package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, W int64, C int64, l []int64, r []int64, p []int64) {

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
    var W int64
    scanner.Scan()
    W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var C int64
    scanner.Scan()
    C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    l := make([]int64, N)
    r := make([]int64, N)
    p := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        l[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        p[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, W, C, l, r, p)
}
