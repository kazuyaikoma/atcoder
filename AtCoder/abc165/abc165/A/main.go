package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(K int64, A int64, B int64) {
	// NOTES: https://scrapbox.io/rustacean/ABC165_Aの数学的解法
	if A/K < B/K || A == A/K*K {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(K, A, B)
}
