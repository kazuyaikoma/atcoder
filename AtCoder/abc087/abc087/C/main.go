package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int64, A [][]int64) {

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
    A := make([][]int64, 2)
    for i := int64(0); i < 2; i++ {
    	A[i] = make([]int64, N)
    }
    for i := int64(0); i < 2; i++ {
        for j := int64(0); j < N; j++ {
            scanner.Scan()
            A[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
        }
    }
	solve(N, A)
}
