package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func cntZero(bin string) int {
	var cnt int
	for _, s := range bin {
		if s == '0' {
			cnt++
		} else {
			return cnt
		}
	}
	return cnt
}

func count(bins []string) int {
	cnt := cntZero(bins[0])
	for _, bin := range bins {
		newCnt := cntZero(bin)
		if newCnt < cnt {
			cnt = newCnt
		}
	}
	return cnt
}

func solve(N int64, A []int64) {
	var bins []string
	for _, a := range A {
		bins = append(bins, reverse(fmt.Sprintf("%b", a)))
	}
	fmt.Println(count(bins))
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
