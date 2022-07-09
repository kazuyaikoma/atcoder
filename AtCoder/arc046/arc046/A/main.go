package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func zorome(n int) bool {
	r := []rune(strconv.Itoa(n))
	x := r[0]
	for _, y := range r {
		if x != y {
			return false
		}
	}
	return true
}

func solve(N int64) {
	for i, cnt := 1, 0; ; i++ {
		if zorome(i) {
			cnt++
			if cnt == int(N) {
				fmt.Println(i)
				return
			}
		}
	}
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
	solve(N)
}
