package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(K int64, A int64, B int64) {
	// A以下のKの最大倍数を求める
	// その数字、もしくはそれに+Kした数字がAからBに収まっているかを調べる
	n := A / K
	if (A <= n*K && n*K <= B) || (A <= (n+1)*K && (n+1)*K <= B) {
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
