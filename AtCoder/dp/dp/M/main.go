package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(N int64, K int64, a []int64) {

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
    a := make([]int64, N)
    for i := int64(0); i < N; i++ {
        scanner.Scan()
        a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
    }
	solve(N, K, a)
}
