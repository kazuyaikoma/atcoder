package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	a = 500
	b = 100
	c = 50
)

func solve(A int64, B int64, C int64, X int64) {
	var cnt int
I:
	for i := 0; i <= int(A); i++ {
		if X < int64(a*i) {
			break I
		}
	J:
		for j := 0; j <= int(B); j++ {
			if X < int64(a*i+b*j) {
				break J
			}
		K:
			for k := 0; k <= int(C); k++ {
				if X < int64(a*i+b*j+c*k) {
					break K
				}
				if int64(a*i+b*j+c*k) == X {
					cnt++
				}
			}
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
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var C int64
	scanner.Scan()
	C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B, C, X)
}
