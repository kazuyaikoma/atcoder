package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(S int64) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var S int64
    scanner.Scan()
    S, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(S)
}
