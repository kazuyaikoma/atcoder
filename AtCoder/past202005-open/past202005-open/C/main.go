package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(A int64, R int64, N int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
    scanner.Scan()
    A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var R int64
    scanner.Scan()
    R, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    var N int64
    scanner.Scan()
    N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, R, N)
}
