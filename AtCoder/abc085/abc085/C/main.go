package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	a = 10000
	b = 5000
	c = 1000
)

func calc(N, Y int) string {
	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {
			k := N - i - j
			if a*i+b*j+c*k == Y {
				return fmt.Sprintf("%d %d %d", i, j, k)
			}
		}
	}
	return fmt.Sprintf("%d %d %d", -1, -1, -1)
}

func solve(N int64, Y int64) {
	s := calc(int(N), int(Y))
	fmt.Println(s)
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
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, Y)
}
