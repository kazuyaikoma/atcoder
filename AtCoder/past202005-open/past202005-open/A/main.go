package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(s string, t string) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var s string
    scanner.Scan()
    s = scanner.Text()
    var t string
    scanner.Scan()
    t = scanner.Text()
	solve(s, t)
}
