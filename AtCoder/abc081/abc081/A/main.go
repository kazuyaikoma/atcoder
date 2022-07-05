package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a string) {
	var cnt = 0
	for _, c := range a {
		if c == '1' {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a string
	scanner.Scan()
	a = scanner.Text()
	solve(a)
}
