package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(c string) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var c string
    scanner.Scan()
    c = scanner.Text()
	solve(c)
}
