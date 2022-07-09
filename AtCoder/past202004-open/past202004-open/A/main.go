package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(S string, T string) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var S string
    scanner.Scan()
    S = scanner.Text()
    var T string
    scanner.Scan()
    T = scanner.Text()
	solve(S, T)
}
