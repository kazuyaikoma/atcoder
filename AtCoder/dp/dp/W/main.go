package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, M int64, l []int64, r []int64, a []int64) {

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
    l := make([]int64, M)
    r := make([]int64, M)
    a := make([]int64, M)
    for i := int64(0); i < M; i++ {
        scanner.Scan()
        l[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        scanner.Scan()
        a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, M, l, r, a)
}
